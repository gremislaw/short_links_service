package util

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func GenerateShortURL(originalURL string) string {
	// Хешируем оригинальный URL с помощью SHA-256
	hash := sha256.Sum256([]byte(originalURL))

	// Кодируем хеш в base64 (для использования допустимых символов)
	encoded := base64.URLEncoding.EncodeToString(hash[:])

	// Используем strings.Builder для замены символов
	var builder strings.Builder
	builder.Grow(10) // Резервируем место для 10 символов

	for i := 0; i < 10; i++ {
		char := encoded[i]
		switch char {
		case ' ':
			builder.WriteByte('_')
		case '+':
			builder.WriteByte('_')
		case '/':
			builder.WriteByte('_')
		case '-':
			builder.WriteByte('_')
		case '=':
			builder.WriteByte('_')
		default:
			builder.WriteByte(char)
		}
	}

	return builder.String()
}
