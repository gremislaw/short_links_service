package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	GRPCAddr    string
	GatewayAddr string
	StorageType string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	return Config{
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		GRPCAddr:    os.Getenv("GRPC_ADDRESS"),
		GatewayAddr: os.Getenv("GATEWAY_ADDRESS"),
		StorageType: os.Getenv("STORAGE_TYPE"),
	}, err
}
