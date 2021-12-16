package main

import (
	"beer-api/internal/logs"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"time"
)

type MyServer struct {
	server *http.Server
}

func NewServer(mux *chi.Mux) *MyServer {
	port := os.Getenv("SERVER_PORT")
	s := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	return &MyServer{s}
}

func (s *MyServer) Run() {
	logs.Log().Fatal(s.server.ListenAndServe())
}
