package server

import (
	"context"
	"net/http"
	"time"
)

type server struct {
	httpServer http.Server
}

func NewServer(host, port string, handler http.Handler) *server {
	return &server{
		httpServer: http.Server{
			Addr:           host + ":" + port,
			MaxHeaderBytes: 1 << 20, //1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			Handler:        handler,
		},
	}
}

func (s *server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
