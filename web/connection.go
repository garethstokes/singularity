package web

import (
  "code.google.com/p/go.net/websocket"
  "github.com/garethstokes/singularity/log"
)

type connection struct {
  ws * websocket.Conn
  send chan string
  server * WebServer
}

func newConnection(s * WebServer) * connection {
  conn := new(connection)
  conn.send = make(chan string, 256)
  conn.server = s

  conn.server.register <- conn
  //defer func() { conn.server.unregister <- conn }()

  return conn
}

func (c * connection) read() {
  for {
    var message [20]byte
    length, err := c.ws.Read(message[:])
    if err != nil {
      log.Errorf("%s", err.Error())
      break
    }
    c.server.broadcast <- string(message[:length])
  }
}

func (c * connection) write() {

  for message := range c.send {
    err := websocket.Message.Send(c.ws, message)
    if err != nil {
      log.Errorf("%s", err.Error())
      break
    }
  }
}

func (c * connection) Close() error {
  return c.ws.Close()
}

func connectionWebSocketHandler(ws * websocket.Conn) {
  c := newConnection(currentWebServer)
  c.ws = ws

  go c.write()
  c.read()
}
