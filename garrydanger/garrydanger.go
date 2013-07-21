package main

import (
  "singularity"
  "singularity/log"
)

type Intelligence int

var intelligence Intelligence

func (s * Intelligence) Tick(args * int, result * int) error {
  log.Info("tick")
  return nil
}

func main() {
  singularity.ServerAddress = "127.0.0.1:4333"
  singularity.ClientAddress = "127.0.0.1:4334"
  singularity.Begin(&intelligence)
}

