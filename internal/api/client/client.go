package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    "link_service/internal/api/grpc/api"
)

func main() {
    // Устанавливаем соединение с сервером
    conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := api.NewUrlShortenerClient(conn)

    // Контекст с таймаутом
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Пример вызова метода CreateShortenedUrl
    createResp, err := client.CreateShortenedUrl(ctx, &api.CreateShortenedUrlRequest{
        OriginalUrl: "https://example.com",
    })
    if err != nil {
        log.Fatalf("could not create shortened URL: %v", err)
    }
    log.Printf("Shortened URL: %s", createResp.GetShortenedUrl())

    // Пример вызова метода GetOriginalURL
    getResp, err := client.GetOriginalURL(ctx, &api.GetOriginalURLRequest{
        ShortenedUrl: createResp.GetShortenedUrl(),
    })
    if err != nil {
        log.Fatalf("could not get original URL: %v", err)
    }
    log.Printf("Original URL: %s", getResp.GetOriginalUrl())
}