run:
	docker-compose up --build -d
	go run cmd/main.go

test:
	go test ./...

docker_up:
	sudo docker compose up -d --build

docker_stop:
	sudo docker compose stop

docker_down:
	sudo docker compose down

format:
	go fmt ./...

clean:
	rm -rf ./bin
	rm -rf ./data