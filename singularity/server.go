package main

import (
  "github.com/garethstokes/singularity"
)

func Server() {
  server := new( singularity.Server )
  server.Start()
}
