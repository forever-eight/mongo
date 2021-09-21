package mongo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// запускает сервер
func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + "8888",
		ReadTimeout:    10 * time.Second,
		Handler:        handler,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //бинарный сдвиг
	}
	return s.httpServer.ListenAndServe()
}

// останавливает работу сервера
func (s Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
