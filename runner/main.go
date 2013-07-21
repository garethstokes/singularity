package main

import (
  "os"
  "singularity/log"
  "singularity/server"
)

func main() {
  if len(os.Args) != 2 {
    log.Error( "please run with 'server' argument" )
    return
  }

  if os.Args[1] == "server" {
    server.Start()
  }
}
