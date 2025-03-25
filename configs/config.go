package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Пакет с конфигурациями.
type Config struct {
	Db     DbConfig
	Adress AdressConfig
}

type DbConfig struct {
	Dsn string
}

type AdressConfig struct {
	OurAddr string
	ApiAddr string
}

func LoadConfigs() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("FATAL: Error loading .env file.")
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Adress: AdressConfig{
			OurAddr: os.Getenv("OUR_ADDR"),
			ApiAddr: os.Getenv("API_ADDR"),
		},
	}
}
