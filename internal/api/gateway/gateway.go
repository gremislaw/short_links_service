package gateway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"link_service/internal/api/grpc/api"
)

func StartRestGateway(grpcAddr string, restAddr string) {
	// Создаем контекст
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Настраиваем HTTP мультиплексор
	mux := runtime.NewServeMux()

	// Регистрируем gRPC-Gateway с использованием конечной точки gRPC-сервера
	opts := []grpc.DialOption{grpc.WithInsecure()} // Используем insecure-подключение для примера
	err := api.RegisterUrlShortenerHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		panic(fmt.Sprintf("Failed to register gRPC Gateway: %v", err))
	}

	// Запускаем
	fmt.Println("Starting REST gateway on", restAddr)
	err = http.ListenAndServe(restAddr, mux)
	if err != nil {
		panic(fmt.Sprintf("Failed to start REST server: %v", err))
	}
}