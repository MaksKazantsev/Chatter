run:
	go build main.go
	docker-compose up
	go run main.go -c config/config.local.yaml

initGRPC:
	protoc --proto_path=pkg/grpc --go_opt=paths=source_relative --go_out=./pkg/grpc --go-grpc_opt=paths=source_relative --go-grpc_out=./pkg/grpc messages.proto

