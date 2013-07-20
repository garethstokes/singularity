package client

import (
  "singularity/log"
  "singularity/server"
)

type Singularity int

var singularity Singularity

func (s * Singularity) Tick(args * int, result * int) error {
  log.Info( "tick" )
  return nil
}

func registerAndListen() {
  server.Register(& singularity)
  server.BindAndListenOn(":4334")
}

func Start() {
  log.Info( "Singularity Test Client" )
  log.Info( "=======================" )

  remoteServer := "localhost:4333"
  client, err := server.Dial(remoteServer)
  if err != nil {
    return
  }

  // Synchronous call
  me := new( server.Host )
  me.Name = "garrydanger"
  me.Address = "localhost"
  me.Port = 4334

  log.Info( "Registering singularity" )
  err = client.Call("Grid.Register", & me, nil)
  if err != nil {
    log.Error(err.Error())
  }

  registerAndListen()
}
