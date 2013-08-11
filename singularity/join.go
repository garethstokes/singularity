package main

import (
  "fmt"
  "os"
  "os/exec"
  "path"
)

func buildpath(executable string) (string, error) {
  //cmdline := fmt.Sprintf("%s/bin/go run %s/src/singularity/%s/garrydanger.go", runtime.GOROOT(), gopath, intelligence)
  p, err := exec.LookPath("go")
  if err != nil {
    fmt.Printf("ERROR: %s\n", err.Error())
    return "", err
  }

  return p, nil
}

func run() {
  intelligence := os.Args[2]
  gopath := os.Getenv("GOPATH")

  wd := fmt.Sprintf("%s/src/singularity/%s", gopath, intelligence)
  wd = path.Clean(wd)

  err := os.Chdir(wd)
  if err != nil {
    fmt.Printf( "ERROR 1: %s\n", err.Error() )
  }

  cmdline1, _ := buildpath("go")
  cmdline2 := fmt.Sprintf("%s.go", intelligence)

  cmd := exec.Command(cmdline1, "run", cmdline2)

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err = cmd.Run()
  if err != nil {
    fmt.Printf( "ERROR 2: %s\n", err.Error() )
  }
}

func Join() {
  run()
}
