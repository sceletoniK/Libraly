package app

import (
	"net/http"

	"github.com/rs/cors"
)

func corsSettings() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
		},
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Authorization"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{},
		Debug:              true,
		MaxAge:             3600,
	})
}
