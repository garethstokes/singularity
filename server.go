/*
  Singularity server 
*/

package singularity

import (
  "net"
  "time"
  "net/http"
  "net/rpc"
	"math/rand"
  "github.com/garethstokes/singularity/log"
  "github.com/garethstokes/singularity/web"
)

var (
  gameAddress = "localhost:4333"
  webAddress = "localhost:8080"
)

type Server struct {
  hostTable * HostTable
  grid Grid
  environment * Environment
  webserver * web.WebServer
}

type Grid struct {
  server * Server
}

func (g * Grid) Register(host * RpcHost, result * int) error {

  err := g.server.hostTable.AddRpcHost(host)
  if err != nil {
    return err
  }

  g.server.environment.AddPlayer(host.Name, "human", 1)

  return nil
}

func (s * Server) AddComputerHost() {
  names := []string {
    "Wesley snipes",
    "Harry potter",
    "Blackadder",
    "Tetris",
    "Atari",
    "Sega",
    "Vega",
    "Flash gordon",
    "Zero cool",
    "Crash overide",
  }

	rand.Seed(int64(time.Now().Nanosecond()))
  name := names[rand.Intn(len(names))]

  log.Info("computer: " + name)

  player := new( MemoryHost )
  player.Name = name // use the working directory

  s.hostTable.AddMovableHostOnly(player)
  s.environment.AddPlayer(player.Name, "ai", 2)
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
  name := host.getName()

  move, err := host.PerformMoveOn(s)
  if err != nil {
    s.hostTable.RemoveHostByName(name)
    s.webserver.Broadcast(toJson("remove", name))
    return
  }

  s.environment.Step(name, move)

  player := s.environment.Entities[name]
  if player == nil {
    s.hostTable.RemoveHostByName(name)
    s.webserver.Broadcast(toJson("remove", name))
  } else {
    s.webserver.Broadcast(toJson("update", player))
  }
}

func (s * Server) Start() {
  log.Info( "Singularity Grid Server" )
  log.Info( "=======================" )

  // web server + socket
  s.webserver = web.NewWebServer()
  go s.webserver.Start()

  s.hostTable = NewHostTable()

  s.environment = NewEnvironment()

  grid := new(Grid)
  grid.server = s
  s.Register(grid)
  go s.BindAndListenOn(gameAddress)

  log.Info( "Addding computer players" )
  for i := 0; i < 1; i++ {
    s.AddComputerHost()
  }

  log.Info( "Entering game loop" )

  c := time.Tick(1000 * time.Millisecond)
  for _ = range c {
    for _, host := range s.hostTable.all {
      go s.tick(host)
    }
  }
}
