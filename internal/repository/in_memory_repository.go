package repository

import (
	"context"
	"errors"
	"sync"
	"time"
)

type inMemoryLinkRepository struct {
	originalToShort sync.Map // Оригинальный URL -> сокращенная ссылка
	shortToOriginal sync.Map // Сокращенная ссылка -> оригинальный URL
}

func NewInMemoryRepository() LinkRepository {
	return &inMemoryLinkRepository{}
}

func (r *inMemoryLinkRepository) CreateShortURL(ctx context.Context, originalURL, shortURL string) error {
	// Сохраняем в обоих мапах
	if !r.ExistsURL(ctx, originalURL) {
		r.originalToShort.Store(originalURL, cacheItem{value: shortURL, expiry: time.Now().Add(10 * time.Minute)})
		r.shortToOriginal.Store(shortURL, originalURL)
	}

	cleanupCache(&r.originalToShort, &r.shortToOriginal)
	clearMapIfMemoryExceeded(&r.originalToShort, &r.shortToOriginal)

	return nil
}

func (r *inMemoryLinkRepository) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	// Получаем оригинальный URL по сокращенной ссылке
	originalURL, exists := r.shortToOriginal.Load(shortURL)
	if !exists {
		return "", errors.New("not found in memory: " + shortURL)
	}
	return originalURL.(string), nil
}

func (r *inMemoryLinkRepository) ExistsURL(ctx context.Context, originalURL string) bool {
	// Проверяем, существует ли уже сокращенная ссылка для этого URL
	_, exists := r.originalToShort.Load(originalURL)
	return exists
}
