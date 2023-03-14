package config

import (
	"log"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
)

type Config struct {
	LogLevel     string
	HTTPPort     string `env:"PORT"`
	DBName       string `env:"DBNAME"`
	DBUser       string `env:"DBUSER"`
	DBPassword   string `env:"DBPASSWORD"`
	DBPort       string `env:"DBPORT"`
	AccessTime   string `env:"ACCESSTIME"`
	RefreshTime  string `env:"REFRESHTIME"`
	Key          string `env:"KEY"`
	RentDuration string `env:"RENTDURATION"` //end of evangelion
}

func NewConfig() *Config {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipFlags: false,
		Files:     []string{"config.env"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".env": aconfigdotenv.New(),
		},
	})
	if err := loader.Load(); err != nil {
		log.Fatal(err)
	}
	return &cfg
}
