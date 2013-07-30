package main

import (
  "github.com/garethstokes/singularity/web"
)

func Web() {
  server := web.NewWebServer()
  server.Start()
}
