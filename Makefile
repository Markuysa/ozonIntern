
inmem:
	go run main.go

postgres:
	go run main.go -db

proto:
	protoc --proto_path=internal/proto --go_out=internal/proto/gen --go_opt=paths=source_relative \
				--grpc-gateway_out=internal/proto/gen --grpc-gateway_opt=paths=source_relative \
    			--go-grpc_out=internal/proto/gen --go-grpc_opt=paths=source_relative \
    			 internal/proto/service.proto

