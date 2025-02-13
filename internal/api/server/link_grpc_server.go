package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"link_service/api/grpc"
	"link_service/internal/service"
	"net"
)

func StartGrpcServer(shutdown context.Context, grpcAddr string, service service.LinkService) {
	// Создаем новый gRPC-сервер
	grpcServer := grpc.NewServer()

	// Создаем экземпляр сервера
	linkServer := NewLinkGrpcServer(service)

	// Регистрируем наш сервер (linkServer) в gRPC-сервере.
	api.RegisterUrlShortenerServer(grpcServer, linkServer)

	// Создаем сетевой listener для прослушивания входящих соединений на указанном адресе (grpcAddr)
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logrus.Fatalf("Failed to listen on port %v: %v", grpcAddr, err)
	}
	logrus.Infof("gRPC server running on port %v...", grpcAddr)

	// Запускаем gRPC-сервер, передавая ему listener.
	// Serve блокирует выполнение, пока сервер не будет остановлен.
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			logrus.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()
	logrus.Info("gRPC server has been successfuly started")

	// Ожидание отмены контекста
	<-shutdown.Done()

	// Грациозное завершение gRPC сервера
	logrus.Info("Shutting down gRPC server gracefully...")
	grpcServer.GracefulStop()
	logrus.Info("gRPC server stopped.")
}

type LinkGrpcServer struct {
	api.UnimplementedUrlShortenerServer
	service service.LinkService
}

func NewLinkGrpcServer(svc service.LinkService) *LinkGrpcServer {
	return &LinkGrpcServer{service: svc}
}

func (s *LinkGrpcServer) CreateShortenedUrl(ctx context.Context, req *api.CreateShortenedUrlRequest) (*api.CreateShortenedUrlResponse, error) {
	logrus.Println("CreateShortenedUrl")
	shortenedURL, err := s.service.CreateShortURL(ctx, req.GetOriginalUrl())
	if err != nil {
		return nil, err
	}
	return &api.CreateShortenedUrlResponse{
		ShortenedUrl: shortenedURL,
	}, nil
}

func (s *LinkGrpcServer) GetOriginalURL(ctx context.Context, req *api.GetOriginalURLRequest) (*api.GetOriginalURLResponse, error) {
	logrus.Println("GetOriginalURL")
	originalURL, err := s.service.GetOriginalURL(ctx, req.GetShortenedUrl())
	if err != nil {
		return nil, err
	}
	return &api.GetOriginalURLResponse{
		OriginalUrl: originalURL,
	}, nil
}
