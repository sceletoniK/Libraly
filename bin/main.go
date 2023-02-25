package main

import (
	"log"

	app "github.com/sceletoniK/App"
	"github.com/sceletoniK/DB"
)

func main() {
	db, err := DB.NewDB("postgres://postgres:scell876@localhost:5432/libraly?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Start(db); err != nil {
		log.Fatal(err)
	}
}
