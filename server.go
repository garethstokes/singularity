package singularity

import (
  "net"
  "time"
  "net/http"
  "net/rpc"
  "github.com/garethstokes/singularity/log"
  "github.com/garethstokes/singularity/web"
)

var (
  gameAddress = "localhost:4333"
  webAddress = "localhost:8080"
)

type Server struct {
  hosts HostTable
  grid Grid
  environment * Environment
  webserver * web.WebServer
}

type Grid struct {
  server * Server
}

func (g * Grid) Register(host * Host, result * int) error {
  for name, _ := range g.server.hosts {
    if name == host.Name {
      log.Infof( "Register Update :: %s", host.Name )
      host.errCount = 0
      g.server.hosts[host.Name] = host
      return nil
    }
  }

  log.Infof( "Register New :: %s", host.Name )
  g.server.hosts[host.Name] = host
  g.server.environment.AddPlayer(host.Name)

  return nil
}

func (s * Server) Register(object interface{}) error {
  if e := rpc.Register(object); e != nil {
    log.Error( "Problem registering" )
    return e
  }

  return nil
}

func (s * Server) BindAndListenOn(address string) error {
  l, e := net.Listen("tcp", address)
  if e != nil {
    log.Errorf("listen error:", e.Error())
    return e
  }

  rpc.HandleHTTP()

  http.Serve(l, nil)
  log.Infof( "Server started: %s", address )

  return nil
}

func (s * Server) Dial(server string) (* rpc.Client, error) {
  log.Info( "Dialing " + server )
  client, err := rpc.DialHTTP("tcp", server)
  if err != nil {
    log.Errorf("dialing: %s", err.Error())
    return nil, err
  }

  log.Info( "Connected to " + server )
  return client, nil
}

func (s * Server) tick(host * Host) {
  if host.client == nil {
    var err error
    host.client, err = s.Dial(host.Address)
    if err != nil {
      if host.errCount == 3 {
        log.Infof( "Removing %s from hosts table", host.Name )
        delete(s.hosts,host.Name)
      } else {
        host.errCount++
      }
      return
    }
  }

  entities := s.environment.Entities

  args := new( TickData )
  args.Player = entities[host.Name]
  args.VisableThings = make( [](* Entity), len(entities) )
  i := 0
  for _, e := range entities {
    args.VisableThings[i] = e
    i++;
  }

  result := new(Move)
  result.Direction = &Vector{0,0}
  result.Action = ACTION_MOVE_STOP

  err := host.client.Call("Intelligence.Tick", args, result)
  if err != nil {

    log.Infof( "Removing %s from hosts table", host.Name )
    delete(s.hosts,host.Name)

    player := s.environment.Entities[host.Name]
    s.webserver.Broadcast(toJson("remove", player))

    return
  }

  s.environment.Step(host.Name, result)

  player := s.environment.Entities[host.Name]
  s.webserver.Broadcast(toJson("update", player))
}

func (s * Server) Start() {
  log.Info( "Singularity Grid Server" )
  log.Info( "=======================" )

  // web server + socket
  s.webserver = web.NewWebServer()
  go s.webserver.Start()

  s.hosts = make( HostTable, 0 )
  s.environment = NewEnvironment()

  grid := new(Grid)
  grid.server = s
  s.Register(grid)
  go s.BindAndListenOn(gameAddress)

  log.Info( "Entering game loop" )

  c := time.Tick(1 * time.Second)
  for _ = range c {
    for _, host := range s.hosts {
      go s.tick(host)
    }
  }
}
