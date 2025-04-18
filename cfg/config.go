package cfg

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	DB     Database
	Server Server
	JWT    JWT
}

type Database struct {
	DBPort     string `env:"DB_PORT"`
	DBHost     string `env:"DB_HOST"`
	DBName     string `env:"DB_NAME"`
	User       string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
}

type Server struct {
	ServerPort string `env:"SERVER_PORT"`
}

type JWT struct {
	SecretKey string `env:"JWT_SECRET"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
