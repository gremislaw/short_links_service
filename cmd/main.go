package main

import (
	"context"
	"link_service/internal/api/gateway"
	"link_service/internal/api/server"
	"link_service/internal/config"
	"link_service/internal/db"
	"link_service/internal/repository"
	"link_service/internal/service"
	"sync"

	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Infof(".env file not found: %v", err)
	}
	logrus.Info("Config has been successfuly loaded")

	// Создание слоя репозитория
	// по параметру определяем тип хранилища
	var repo repository.LinkRepository
	switch cfg.StorageType {
	case "in_memory":
		repo = repository.NewInMemoryRepository()
		logrus.Info("In memory storage has been chosen")
	case "postgres":
		// Подключение к БД
		DB, err := db.NewPostgresDB(cfg)
		if err != nil {
			logrus.Fatalf("Failed to connect to DB: %v", err)
		}
		logrus.Info("Database has been successfuly connected")
		repo = repository.NewPostgresRepository(DB)
		logrus.Info("Postgres storage has been chosen")
	default:
		repo = repository.NewInMemoryRepository()
		logrus.Info("Wrong type, in memory storage has been automatically chosen")
	}

	// Создание слоя сервиса
	service := service.NewLinkService(repo)

	// Контекст с отменой для управления жизненным циклом горутин
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запуск GRPC сервера
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.StartGrpcServer(ctx, cfg.GRPCAddr, service)
	}()

	// Запуск REST Gateway
	wg.Add(1)
	go func() {
		defer wg.Done()
		gateway.StartRestGateway(ctx, cfg.GRPCAddr, cfg.GatewayAddr)
	}()

	// Канал для обработки сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала завершения
	<-stop
	logrus.Info("Received shutdown signal. Gracefully shutting down...")

	// Отмена контекста для остановки всех горутин
	cancel()

	// Ожидание завершения всех горутин
	wg.Wait()
	logrus.Info("All servers stopped. Exiting.")
}
