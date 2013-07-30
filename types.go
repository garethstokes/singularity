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
  Name string         `json:"name"`
  Position * Vector   `json:"position"`
  Direction * Vector  `json:"direction"`
  Speed int           `json:"-"`
  Action MoveAction   `json:"-"`
}

type HostTable map[string] * Host

type MoveAction int

type Move struct {
  Direction * Vector
  Action MoveAction
}
