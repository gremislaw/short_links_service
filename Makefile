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

generate_sqlc:
	sqlc generate

generate_grpc:
	protoc -I proto \
  -I $GOPATH/pkg/mod/google.golang.org/genproto/googleapis \
  --go_out=internal/api/grpc --go-grpc_out=internal/api/grpc --grpc-gateway_out=internal/api/grpc \
  --plugin=protoc-gen-grpc-gateway=$(which protoc-gen-grpc-gateway) \
  --plugin=protoc-gen-go=$(which protoc-gen-go) \
  --plugin=protoc-gen-go-grpc=$(which protoc-gen-go-grpc) \
  proto/main.proto


clean:
	rm -rf ./bin
	rm -rf ./data