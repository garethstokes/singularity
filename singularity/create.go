package main

import (
  "os"
  "fmt"
  "path"
  "io/ioutil"
)

var (
  template = `
package main

import (
  "github.com/garethstokes/singularity"
)

type Intelligence int

var intelligence Intelligence

func (s * Intelligence) Tick(args * int, result * int) error {
  // Start hacking here

  return nil
}

func main() {
  singularity.ServerAddress = "127.0.0.1:4333"
  singularity.ClientAddress = "127.0.0.1:4334"
  singularity.Begin(&intelligence)
}
`
)

func Create() {
  if len(os.Args) < 3 {
    fail()
    return
  }

  handle := os.Args[2]

  fmt.Printf( "initializing new intelligence with handle: %s.\n", handle )

  installpath := fmt.Sprintf( "%s/src/singularity/%s", os.Getenv("GOPATH"), handle)
  installpath = path.Clean(installpath)

  err := os.MkdirAll(installpath, 0777)
  if err != nil {
    fmt.Printf( "ERROR: %s\n", err.Error() )
    return
  }

  filepath := fmt.Sprintf( "%s/%s.go", installpath, handle )

  err = ioutil.WriteFile( filepath, []byte(template), 0644 )
  if err != nil {
    fmt.Printf( "ERROR: %s\n", err.Error() )
    return
  }

  fmt.Printf( "Intellegence created at path: %s\n", filepath )
}
