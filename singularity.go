/*
  Singularity client 
*/

package singularity

import (
  "github.com/garethstokes/singularity/log"
  "os"
  "path"
  "encoding/json"
)

var (
  ServerAddress = "localhost:4333"
  ClientAddress = "localhost:4334"
)

const (
  _                                   = iota
  ACTION_MOVE_FORWARD    MoveAction   = (10 * iota)
  ACTION_MOVE_BACKWARD
  ACTION_MOVE_STOP
  ACTION_MOVE_TURN
)

func toJson(key string, item interface{}) string {
  kv := make(map[string]interface{})
  kv[key] = item
	b, err := json.Marshal(kv)
	if err != nil {
    log.Errorf(err.Error())
	}
	return string(b);
}

func Begin(intelligence interface{}) {
  log.Info( "Singularity Client" )
  log.Info( "==================" )

  server := new( Server )
  client, err := server.Dial(ServerAddress)
  if err != nil {
    return
  }

  wd, _ := os.Getwd()

  me := new( RpcHost )
  me.Name = path.Base(wd) // use the working directory
  me.Address = ClientAddress

  log.Infof("Registering singularity for %s@%s", me.Name, ClientAddress)
  err = client.Call("Grid.Register", & me, nil)
  if err != nil {
    log.Error(err.Error())
  }


  server.Register(intelligence)
  server.BindAndListenOn(ClientAddress)
}
