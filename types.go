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
  Position * Point
  Entities []Entity
}

type Entity struct {
  Name string
  Position * Point
}

type Point struct {
  X int
  Y int
}

type HostTable map[string] * Host
type Grid int
type Move int
