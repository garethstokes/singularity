package singularity

import (
  "github.com/garethstokes/singularity/log"
  "os"
  "path"
)

var (
  ServerAddress = "localhost:4333"
  ClientAddress = "localhost:4334"
)

func Begin(intelligence interface{}) {
  log.Info( "Singularity Test Client" )
  log.Info( "=======================" )

  server := new( Server )
  client, err := server.Dial(ServerAddress)
  if err != nil {
    return
  }

  wd, _ := os.Getwd()

  me := new( Host )
  me.Name = path.Base(wd) // use the working directory
  me.Address = ClientAddress

  log.Infof("Registering singularity for %s@%s", me.Name, ClientAddress)
  err = client.Call("Grid.Register", & me, nil)
  if err != nil {
    log.Error(err.Error())
  }


  server.Register(intelligence)
  server.BindAndListenOn(":4334")
}
