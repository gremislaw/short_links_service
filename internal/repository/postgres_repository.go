package repository

import (
	"context"
	"link_service/internal/db"
)

type postgresRepository struct {
	queries *db.Queries
}

func NewPostgresRepository(queries *db.Queries) LinkRepository {
	return &postgresRepository{queries: queries}
}

func (r *postgresRepository) CreateShortURL(ctx context.Context, originalURL, shortenedURL string) error {
	// Вставка сокращённого URL в базу данных
	if !r.ExistsURL(ctx, originalURL) {
		return r.queries.CreateShortURL(ctx, db.CreateShortURLParams{
			OriginalUrl:  originalURL,
			ShortenedUrl: shortenedURL,
		})
	}
	return nil
}

func (r *postgresRepository) GetOriginalURL(ctx context.Context, shortenedURL string) (string, error) {
	// Получаем оригинальный URL по сокращённому URL
	return r.queries.GetOriginalURL(ctx, shortenedURL)
}

func (r *postgresRepository) ExistsURL(ctx context.Context, originalURL string) bool {
	// Проверяем, существует ли уже сокращенная ссылка для этого URL
	str, _ := r.queries.ExistsURL(ctx, originalURL)
	return str != ""
}
