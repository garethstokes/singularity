package main

import (
  "os"
  "fmt"
)

var message = `
  -={ the singularity is near }=-

  a new digital entity has arrived and we now live 
  on a planet in which humans are no longer
  the superior intellect. we must evolve, adapt 
  to this new digital environment and prove 
  once and for all, that humanity is worth saving.

  the commands to enter the game are as follows...

  create:
  singularity create <player name>

  server:
  singularity server

  join:
  singularity join <player name>

`

func fail() {
  fmt.Print(message)
}

func main() {
  if len(os.Args) < 2 {
    fail()
    return
  }

  if os.Args[1] == "create" {
    Create()
  }

  if os.Args[1] == "server" {
    Server()
  }

  if os.Args[1] == "join" {
    Join()
  }

  if os.Args[1] == "web" {
    Web()
  }
}
