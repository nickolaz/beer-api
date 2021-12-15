package main

import (
	"beer-api/internal/logs"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

type MyServer struct {
	server *http.Server
}

func NewServer(mux *chi.Mux) *MyServer {
	s := &http.Server{
		Addr:              ":9000",
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
