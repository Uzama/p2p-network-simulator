package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (s *HTTPServer) Start() error {
	r := initRouter()

	address := fmt.Sprintf("0.0.0.0:%d", 8080)

	server := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
		Handler:      r,
	}

	s.server = server

	go s.listen()

	log.Println("server started at 0.0.0.0:8080")

	return nil
}

func (s HTTPServer) listen() {
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func (s HTTPServer) Shutdown(ctx context.Context) {
	s.server.SetKeepAlivesEnabled(false)

	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
