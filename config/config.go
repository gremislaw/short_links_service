package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	GRPCAddr    string
	GatewayAddr string
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return Config{}, err
	}
	return Config{
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		GRPCAddr:    os.Getenv("GRPC_ADDRESS"),
		GatewayAddr: os.Getenv("GATEWAY_ADDRESS"),
	}, nil
}
