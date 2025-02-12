// internal/api/grpc/link_grpc_server.go
package server

import (
	"log"
	"context"
	"link_service/internal/api/grpc/api"
	"link_service/internal/service"
)

type LinkGrpcServer struct {
	api.UnimplementedUrlShortenerServer
	service service.LinkService
}

func NewLinkGrpcServer(svc service.LinkService) *LinkGrpcServer {
	return &LinkGrpcServer{service: svc}
}

func (s *LinkGrpcServer) CreateShortenedUrl(ctx context.Context, req *api.CreateShortenedUrlRequest) (*api.CreateShortenedUrlResponse, error) {
	log.Println("CreateShortenedUrl")
	shortenedURL, err := s.service.CreateShortURL(ctx, req.GetOriginalUrl())
	if err != nil {
		return nil, err
	}
	return &api.CreateShortenedUrlResponse{
		ShortenedUrl: shortenedURL,
	}, nil
}

func (s *LinkGrpcServer) GetOriginalURL(ctx context.Context, req *api.GetOriginalURLRequest) (*api.GetOriginalURLResponse, error) {
	log.Println("GetOriginalURL")
	originalURL, err := s.service.GetOriginalURL(ctx, req.GetShortenedUrl())
	if err != nil {
		return nil, err
	}
	return &api.GetOriginalURLResponse{
		OriginalUrl: originalURL,
	}, nil
}
