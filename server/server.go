package server

import (
  "os"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "ui-mockup-backend"
)

type Server struct {
  router *mux.Router
  config *root.ServerConfig
}

func NewServer(stdService root.StandardService, userService root.UserService, config *root.Config) *Server {
  s := Server { 
    router: mux.NewRouter(),
    config: config.Server }
  
  a := authHelper{config.Auth.Secret}
  NewStandardRouter(stdService, s.getSubrouter("/standard"), &a)
  NewUserRouter(userService, s.getSubrouter("/user"), &a)
  return &s
}

func(s *Server) Start() {
  log.Println("Listening on port " + s.config.Port)
  if err := http.ListenAndServe(s.config.Port, handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
      log.Fatal("http.ListenAndServe: ", err)
  }
}

func(s *Server) getSubrouter(path string) *mux.Router {
  return s.router.PathPrefix(path).Subrouter()
}
