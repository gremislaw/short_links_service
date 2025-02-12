package util

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func GenerateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	shortenedURL := make([]byte, 10)
	for i := range shortenedURL {
		shortenedURL[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortenedURL)
}
