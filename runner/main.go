package main

import (
  "os"
  "singularity/log"
  "singularity/server"
  "singularity/client"
)

func main() {
  if len(os.Args) != 2 {
    log.Error( "please run with either a 'server' or 'client' argument" )
    return
  }

  if os.Args[1] == "server" {
    server.Start()

  }

  if os.Args[1] == "client" {
    client.Start()
  }
}
