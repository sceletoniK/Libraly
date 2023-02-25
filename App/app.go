package app

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

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
	InitLogger()
	s.logger.Info("Server start")
	return s.httpServer.Serve(listener)
}
