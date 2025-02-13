package service_test

import (
	"context"
	"link_service/internal/service"
	"link_service/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Мокаем репозиторий
type MockLinkRepository struct {
	mock.Mock
}

func (m *MockLinkRepository) CreateShortURL(ctx context.Context, originalURL, shortenedURL string) error {
	args := m.Called(ctx, originalURL, shortenedURL)
	return args.Error(0)
}

func (m *MockLinkRepository) GetOriginalURL(ctx context.Context, shortenedURL string) (string, error) {
	args := m.Called(ctx, shortenedURL)
	return args.String(0), args.Error(1)
}

func TestCreateShortURL(t *testing.T) {
	mockRepo := new(MockLinkRepository)
	linkService := service.NewLinkService(mockRepo)

	// Даем моку ожидания
	originalURL := "http://example.com"
	shortenedURL := util.GenerateShortURL(originalURL)
	mockRepo.On("CreateShortURL", mock.Anything, originalURL, shortenedURL).Return(nil)

	// Генерируем сокращенную ссылку с помощью сервиса
	result, err := linkService.CreateShortURL(context.Background(), originalURL)

	// Проверяем результаты
	assert.NoError(t, err)
	assert.Equal(t, shortenedURL, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateShortURL_Error(t *testing.T) {
	mockRepo := new(MockLinkRepository)
	linkService := service.NewLinkService(mockRepo)

	originalURL := "http://example.com"
	shortenedURL := util.GenerateShortURL(originalURL)
	mockRepo.On("CreateShortURL", mock.Anything, originalURL, shortenedURL).Return(assert.AnError)

	// Генерируем сокращенную ссылку с помощью сервиса
	result, err := linkService.CreateShortURL(context.Background(), originalURL)

	// Проверяем, что ошибка возвращена
	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetOriginalURL(t *testing.T) {
	mockRepo := new(MockLinkRepository)
	linkService := service.NewLinkService(mockRepo)

	originalURL := "http://example.com"
	shortenedURL := util.GenerateShortURL(originalURL)
	mockRepo.On("GetOriginalURL", mock.Anything, shortenedURL).Return(originalURL, nil)

	// Получаем оригинальную ссылку через сервис
	result, err := linkService.GetOriginalURL(context.Background(), shortenedURL)

	// Проверяем результат
	assert.NoError(t, err)
	assert.Equal(t, originalURL, result)
	mockRepo.AssertExpectations(t)
}

func TestGetOriginalURL_Error(t *testing.T) {
	mockRepo := new(MockLinkRepository)
	linkService := service.NewLinkService(mockRepo)

	shortenedURL := "http://short.ly/abcd"
	mockRepo.On("GetOriginalURL", mock.Anything, shortenedURL).Return("", assert.AnError)

	// Получаем оригинальную ссылку через сервис
	result, err := linkService.GetOriginalURL(context.Background(), shortenedURL)

	// Проверяем, что ошибка возвращена
	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}
