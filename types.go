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

type HostTable map[string] * Host
type Grid int
