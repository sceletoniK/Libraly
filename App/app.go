package app

import (
	"net"
	"net/http"
	"time"
)

func Start(l Libraly) error {
	s := NewServer(l)
	handler := http.Handler(s.router)
	s.httpServer = &http.Server{
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	return s.httpServer.Serve(listener)
}
