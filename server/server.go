package server

import (
  "net"
  "time"
  "net/http"
  "net/rpc"
  "github.com/garethstokes/singularity/log"
)

type Host struct {
  Name string
  Address string
  client * rpc.Client
  errCount int
}

type HostTable map[string] * Host
type Grid int

var (
  hosts HostTable
  grid Grid
)

func (s * Grid) Register(host * Host, result * int) error {
  for name, _ := range hosts {
    if name == host.Name {
      log.Infof( "Register Update :: %s", host.Name )
      host.errCount = 0
      hosts[host.Name] = host
      return nil
    }
  }

  log.Infof( "Register New :: %s", host.Name )
  //hosts = append( hosts, * host )
  hosts[host.Name] = host

  return nil
}

func Register(object interface{}) error {
  if e := rpc.Register(object); e != nil {
    log.Error( "Problem registering" )
    return e
  }

  return nil
}

func BindAndListenOn(address string) error {
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

func Dial(server string) (* rpc.Client, error) {
  log.Info( "Dialing " + server )
  client, err := rpc.DialHTTP("tcp", server)
  if err != nil {
    log.Errorf("dialing: %s", err.Error())
    return nil, err
  }

  log.Info( "Connected to " + server )
  return client, nil
}

func tick(host * Host) {
  if host.client == nil {
    var err error
    host.client, err = Dial(host.Address)
    if err != nil {
      if host.errCount == 3 {
        log.Infof( "Removing %s from hosts table", host.Name )
        delete(hosts,host.Name)
      } else {
        host.errCount++
      }
      return
    }
  }

  args := 1
  result := 1
  err := host.client.Call("Intelligence.Tick", &args, &result)
  if err != nil{
    log.Infof( "Removing %s from hosts table", host.Name )
    delete(hosts,host.Name)
  }
}

func Start() {
  log.Info( "Singularity Grid Server" )
  log.Info( "=======================" )

  hosts = make( HostTable, 0 )

  Register(& grid)
  go BindAndListenOn(":4333")

  log.Info( "Entering game loop" )

  c := time.Tick(1 * time.Second)
  for _ = range c {
    for _, host := range hosts {
      go tick(host)
    }
  }
}
