package main

import (
  "os"
  "fmt"
)

func fail() {
  fmt.Print("the singularity has failed.\n")
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

}
