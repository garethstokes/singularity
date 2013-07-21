package singularity

import (
  "singularity/log"
  "singularity/server"
  "os"
  "path"
)

var (
  ServerAddress = "localhost:4333"
  ClientPort = 4334
  ClientAddress = "localhost"
)

func Begin(intelligence interface{}) {
  log.Info( "Singularity Test Client" )
  log.Info( "=======================" )

  client, err := server.Dial(ServerAddress)
  if err != nil {
    return
  }

  wd, _ := os.Getwd()

  me := new( server.Host )
  me.Name = path.Base(wd) // use the working directory
  me.Address = ClientAddress

  log.Infof("Registering singularity for %s:%d", ClientAddress, ClientPort)
  err = client.Call("Grid.Register", & me, nil)
  if err != nil {
    log.Error(err.Error())
  }


  server.Register(intelligence)
  server.BindAndListenOn(":4334")
}
