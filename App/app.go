package app

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/sceletoniK/config"
	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func Start(l Libraly, cfg *config.Config) error {
	s := NewServer(l)
	handler := http.Handler(s.router)
	s.httpServer = &http.Server{
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	s.logger.Info(cfg.HTTPPort)
	listener, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		return err
	}
	InitLogger()
	s.logger.Info("Server start")
	return s.httpServer.Serve(listener)
}
