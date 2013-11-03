/*
  Singularity server 
*/

package singularity

import (
  "net"
  "time"
  "net/http"
  "net/rpc"
  "errors"
  "github.com/garethstokes/singularity/log"
  "github.com/garethstokes/singularity/web"
)

var (
  gameAddress = "localhost:4333"
  webAddress = "localhost:8080"
)

type Server struct {
  hosts HostTable
  rpcHosts map[string] * RpcHost
  grid Grid
  environment * Environment
  webserver * web.WebServer
}

type Grid struct {
  server * Server
}

func (g * Grid) Register(host * RpcHost, result * int) error {
  for name, h := range g.server.rpcHosts {

    // do a simple check if someone else is already using 
    // that port
    if host.Address == h.Address {
      return errors.New("ClientAddress is already in use.")
    }

    // maybe the user is already registered and just wants
    // to update their information?
    if name == host.Name {
      log.Infof( "Register Update :: %s", host.Name )
      host.resetErrors()

      g.server.hosts[host.Name] = host
      g.server.rpcHosts[host.Name] = host

      return nil
    }
  }

  log.Infof( "Register New :: %s", host.Name )

  g.server.hosts[host.Name] = host
  g.server.rpcHosts[host.Name] = host

  g.server.environment.AddPlayer(host.Name)

  return nil
}

func (s * Server) AddComputerHost() {
  name := "Wesley snipes"
  player := new( MemoryHost )
  player.Name = name // use the working directory

  s.hosts[name] = player
  s.environment.AddPlayer(player.Name)
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

func (s * Server) tick(host Movable) {
  move, err := host.PerformMoveOn(s)
  if err != nil {
    return
  }

  s.environment.Step(host.getName(), move)

  player := s.environment.Entities[host.getName()]
  s.webserver.Broadcast(toJson("update", player))
}

func (s * Server) Start() {
  log.Info( "Singularity Grid Server" )
  log.Info( "=======================" )

  // web server + socket
  s.webserver = web.NewWebServer()
  go s.webserver.Start()

  s.hosts = make( HostTable, 0 )
  s.rpcHosts = make( map[string] * RpcHost, 0 )

  s.environment = NewEnvironment()

  grid := new(Grid)
  grid.server = s
  s.Register(grid)
  go s.BindAndListenOn(gameAddress)

  log.Info( "Addding computer players" )
  s.AddComputerHost()

  log.Info( "Entering game loop" )

  c := time.Tick(500 * time.Millisecond)
  for _ = range c {
    for _, host := range s.hosts {
      go s.tick(host)
    }
  }
}
