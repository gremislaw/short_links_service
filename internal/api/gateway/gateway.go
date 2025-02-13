package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"link_service/internal/api/grpc/api"
	"net/http"
)

func StartRestGateway(shutdown context.Context, grpcAddr string, restAddr string) {
	// Создаем контекст
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Настраиваем HTTP мультиплексор
	mux := runtime.NewServeMux()

	// Регистрируем gRPC-Gateway с использованием конечной точки gRPC-сервера
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())} // Используем insecure-подключение для примера
	err := api.RegisterUrlShortenerHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		logrus.Errorf("Failed to register gRPC Gateway: %v", err)
		return
	}

	// Создаем HTTP сервер
	httpServer := &http.Server{
		Addr:    restAddr,
		Handler: mux,
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		logrus.Infof("Starting REST gateway on %s", restAddr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("Failed to start REST server: %v", err)
		}
	}()
	logrus.Info("REST gateway has been successfuly started")

	// Грациозное завершение HTTP сервера
	<-shutdown.Done()

	logrus.Info("Shutting down REST gateway gracefully...")
	if err := httpServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Failed to shutdown REST gateway: %v", err)
	} else {
		logrus.Info("REST gateway has been successfully stopped")
	}
}
