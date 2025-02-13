package repository

import "context"

type LinkRepository interface {
	CreateShortURL(ctx context.Context, originalURL, shortenedURL string) error
	GetOriginalURL(ctx context.Context, shortenedURL string) (string, error)
}
