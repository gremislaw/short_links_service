package repository

import (
	"context"
	"link_service/internal/db"
)

type LinkRepository interface {
	CreateShortURL(ctx context.Context, originalURL, shortenedURL string) error
	GetOriginalURL(ctx context.Context, shortenedURL string) (string, error)
}

type linkRepository struct {
	queries *db.Queries
}

func NewLinkRepository(queries *db.Queries) LinkRepository {
	return &linkRepository{queries: queries}
}

func (r *linkRepository) CreateShortURL(ctx context.Context, originalURL, shortenedURL string) error {
	// Вставка сокращённого URL в базу данных
	return r.queries.CreateShortURL(ctx, db.CreateShortURLParams{
		OriginalUrl:  originalURL,
		ShortenedUrl: shortenedURL,
	})
}

func (r *linkRepository) GetOriginalURL(ctx context.Context, shortenedURL string) (string, error) {
	// Получаем оригинальный URL по сокращённому URL
	return r.queries.GetOriginalURL(ctx, shortenedURL)
}
