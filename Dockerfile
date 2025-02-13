#Docker Pipeline
FROM golang:1.23.2-alpine3.20 as builder
WORKDIR /app
COPY . /app
# Создаем директорию bin, если она не существует
RUN go mod download && \
    go build -o ./bin/short_links_service ./cmd

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/bin/short_links_service ./bin/short_links_service
EXPOSE 9000
ENTRYPOINT ["./bin/short_links_service"]
