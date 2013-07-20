package log

import (
  "fmt"
  "time"
)

func Info( message string ) {
  fmt.Printf( "\033[37;1m%s: %s\033[0m\n", time.Now().Format(time.StampMilli), message )
}

func Infof(message string, args ...interface{}) {
  text := fmt.Sprintf( message, args... )
  fmt.Printf( "\033[37;1m%s: %s\033[0m\n", time.Now().Format(time.StampMilli), text )
}

func Error( message string ) {
  // the commented out one is green
  //fmt.Printf( "\033[32;1m%s: %s \033[0m \n", time.Now().Format(time.StampMilli), message )
  fmt.Printf( "\033[31;1m%s ERROR: %s\033[0m \n", time.Now().Format(time.StampMilli), message )
}

func Errorf(message string, args ...interface{}) {
  text := fmt.Sprintf( message, args... )
  fmt.Printf( "\033[31;1m%s: %s\033[0m\n", time.Now().Format(time.StampMilli), text )
}
