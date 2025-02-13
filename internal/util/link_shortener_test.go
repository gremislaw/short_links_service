package util_test

import (
	"testing"
	"link_service/internal/util"
	"github.com/stretchr/testify/assert"
	"regexp"
)

func TestGenerateShortURL(t *testing.T) {
	tests := []struct {
		originalURL string
	}{
		{originalURL: "http://example.com"},
		{originalURL: "https://google.com"},
		{originalURL: "http://ozon.ru"},
		{originalURL: "https://github.com"},
		{originalURL: "https://stackoverflow.com"},
	}

	for _, tt := range tests {
		t.Run(tt.originalURL, func(t *testing.T) {
			// Генерируем сокращенную ссылку
			shortURL := util.GenerateShortURL(tt.originalURL)

			// Проверяем длину
			assert.Len(t, shortURL, 10, "Generated short URL should have exactly 10 characters")

			// Проверяем уникальность (сравниваем с другой ссылкой для этого же URL)
			shortURL2 := util.GenerateShortURL(tt.originalURL)
			assert.Equal(t, shortURL, shortURL2, "Generated short URLs should be the same for the same original URL")

			// Проверяем, что ссылка состоит только из допустимых символов
			validPattern := "^[a-zA-Z0-9_]{10}$"
			matched, err := regexp.MatchString(validPattern, shortURL)
			if err != nil {
				t.Fatalf("Error during regex match: %v", err)
			}

			assert.True(t, matched, "Generated short URL contains invalid characters")
		})
	}
}

func TestGenerateShortURL_UniqueForDifferentURLs(t *testing.T) {
	tests := []struct {
		originalURL1 string
		originalURL2 string
	}{
		{"http://example.com", "https://google.com"},
		{"http://ozon.ru", "https://github.com"},
		{"https://github.com", "https://stackoverflow.com"},
	}

	for _, tt := range tests {
		t.Run(tt.originalURL1+" vs "+tt.originalURL2, func(t *testing.T) {
			// Генерируем сокращенные ссылки для разных URL
			shortURL1 := util.GenerateShortURL(tt.originalURL1)
			shortURL2 := util.GenerateShortURL(tt.originalURL2)

			// Проверяем, что ссылки разные для разных URL
			assert.NotEqual(t, shortURL1, shortURL2, "Generated short URLs should be different for different original URLs")
		})
	}
}
