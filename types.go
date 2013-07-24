package singularity

import (
  "net/rpc"
)

type Host struct {
  Name string
  Address string
  client * rpc.Client
  errCount int
}

type TickData struct {
  Player * Entity
  VisableThings [](* Entity)
}

type Entity struct {
  Name string
  Position * Vector
  Direction * Vector
  Speed int
  Action MoveAction
}

type HostTable map[string] * Host

type MoveAction int

type Move struct {
  Direction * Vector
  Action MoveAction
}
