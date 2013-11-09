package singularity

import (
  "net/rpc"
  //"github.com/garethstokes/singularity/log"
)

type RpcHost struct {
  Host
  Address string
  client * rpc.Client
}

func (host * RpcHost) getName() string {
  return host.Name
}

func (host * RpcHost) resetErrors() {
  host.errCount = 0
}

func (host * RpcHost) PerformMoveOn(s * Server) (* Move, error) {
  if host.client == nil {
    var err error
    host.client, err = s.Dial(host.Address)
    if err != nil {
      if host.errCount == 3 {
        return nil, err
      }
      return nil, nil
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

    player := s.environment.Entities[host.Name]
    s.webserver.Broadcast(toJson("remove", player.Name))

    return nil, err
  }

  return result, nil
}

