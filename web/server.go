package web

import (
  "os"
  "path"
  "net/http"
  "code.google.com/p/go.net/websocket"
  "github.com/garethstokes/singularity/log"
)

const (
  address = "localhost:8080"
)

var (
  currentWebServer * WebServer
)

type WebServer struct {

  // connections
  connections map[* connection]bool

  // from the connections
  broadcast chan string

  // requests from the clients
  register chan * connection

  // request from the clients
  unregister chan * connection
}

func NewWebServer() * WebServer {
  ws := &WebServer{
    connections: make(map[* connection]bool),
    broadcast: make(chan string),
    register: make(chan * connection),
    unregister: make(chan * connection),
  }

  currentWebServer = ws
  return ws
}

func (w * WebServer) Start() {
  log.Infof( "Web server running on: %s", address )

  // route the web socket events
  go func() {
    for {
        select {
        case c := <-w.register:
          log.Info("registering web socket client")
          w.connections[c] = true
        case c := <-w.unregister:
          log.Info("unregistering web socket client")
          delete(w.connections, c)
          close(c.send)
        case m := <-w.broadcast:
          for c := range w.connections {
              select {
                case c.send <- m:
                default:
                  delete(w.connections, c)
                  close(c.send)
                  go c.Close()
              }
          }
        }
    }
  }()

  gopath := os.Getenv("GOPATH")
  basepath := gopath + "/src/github.com/garethstokes/singularity/web/assets/"

  http.Handle("/ws", websocket.Handler(connectionWebSocketHandler))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    p := r.URL.Path

    if len(p) >= 4 && p[:4] == "/js/" {
      http.ServeFile(w, r, path.Join(basepath + p))
      return
    }

    if len(p) >= 13 && p[:13] == "/stylesheets/" {
      http.ServeFile(w, r, path.Join(basepath + p))
      return
    }

    if len(p) >= 8 && p[:8] == "/images/" {
      http.ServeFile(w, r, path.Join(basepath + p))
      return
    }

    if len(p) >= 7 && p[:7] == "/fonts/" {
      http.ServeFile(w, r, path.Join(basepath + p))
      return
    }

    var filename = basepath + "index.html"
    filename = path.Clean(filename)
    http.ServeFile(w, r, filename)
  })

  err := http.ListenAndServe(address, nil)
  if err != nil {
    log.Errorf("ERROR: %s", err.Error())
  }
}

func (w * WebServer) Broadcast(message string) {
  w.broadcast <- message
}
