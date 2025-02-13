run:
	sudo docker compose up --build -d

install: install_tools
	go mod tidy

install_tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2

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
  --go_out=api/grpc --go-grpc_out=api/grpc --grpc-gateway_out=api/grpc \
  --plugin=protoc-gen-grpc-gateway=$(which protoc-gen-grpc-gateway) \
  --plugin=protoc-gen-go=$(which protoc-gen-go) \
  --plugin=protoc-gen-go-grpc=$(which protoc-gen-go-grpc) \
  api/proto/main.proto


clean:
	rm -rf ./bin
	rm -rf ./data