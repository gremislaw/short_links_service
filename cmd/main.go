package main

import (
	"google.golang.org/grpc"
	"link_service/config"
	"link_service/internal/api/grpc/api"
	"link_service/internal/api/server"
	"link_service/internal/api/gateway"
	"link_service/internal/db"
	"link_service/internal/repository"
	"link_service/internal/service"
	"log"
	"net"
	"fmt"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	DB, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	repo := repository.NewLinkRepository(DB)
	service := service.NewLinkService(repo)

	grpcServer := grpc.NewServer()
	linkServer := server.NewLinkGrpcServer(service)
	api.RegisterUrlShortenerServer(grpcServer, linkServer)

	// Запуск gRPC сервера в горутине
	go func() {
		listener, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatalf("Failed to listen on port 8081: %v", err)
		}
		fmt.Println("gRPC server running on port 8081...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Запуск REST Gateway на порту 8080
	gateway.StartRestGateway(":8081", ":8080")
}
