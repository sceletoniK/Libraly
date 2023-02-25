package main

import (
	"log"

	app "github.com/sceletoniK/App"
	"github.com/sceletoniK/DB"
	"github.com/sceletoniK/config"
)

func main() {
	cfg := config.NewConfig()
	db, err := DB.NewDB("postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@localhost" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Start(db, cfg); err != nil {
		log.Fatal(err)
	}
}
