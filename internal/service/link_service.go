package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"link_service/internal/repository"
	"link_service/internal/util"
)

type LinkService interface {
	CreateShortURL(ctx context.Context, originalURL string) (string, error)
	GetOriginalURL(ctx context.Context, shortenedURL string) (string, error)
}

type linkService struct {
	repository repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) LinkService {
	return &linkService{
		repository: repo,
	}
}

func (s *linkService) CreateShortURL(ctx context.Context, originalURL string) (string, error) {
	logrus.Infof("CreateShortURL got: %v", originalURL)
	shortenedURL := util.GenerateShortURL(originalURL)

	// Сохраняем сокращённую ссылку в базе данных
	err := s.repository.CreateShortURL(ctx, originalURL, shortenedURL)
	if err != nil {
		logrus.Errorf("Error when saving short url: %v", err)
		return "", err
	}

	logrus.Infof("CreateShortURL sent: %v", shortenedURL)
	return shortenedURL, nil
}

func (s *linkService) GetOriginalURL(ctx context.Context, shortenedURL string) (string, error) {
	logrus.Infof("GetOriginalURL got: %v", shortenedURL)
	originalURL, err := s.repository.GetOriginalURL(ctx, shortenedURL)
	if err != nil {
		logrus.Errorf("Error when getting original url: %v", err)
		return "", err
	}

	logrus.Infof("GetOriginalURL sent: %v", originalURL)
	return originalURL, nil
}
