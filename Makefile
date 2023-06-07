
inmem:
	go run app/cmd/main.go

postgres:
	go run app/cmd/main.go -db

proto:
	protoc --proto_path=internal/proto --go_out=internal/proto/gen --go_opt=paths=source_relative \
				--grpc-gateway_out=internal/proto/gen --grpc-gateway_opt=paths=source_relative \
    			--go-grpc_out=internal/proto/gen --go-grpc_opt=paths=source_relative \
    			 internal/proto/service.proto

cover:
	go test ./... -coverprofile=cover.out
	go tool cover -html=cover.out
	rm cover.out

gen:
	mockgen --source=internal/database/database.go \
	--destination=internal/mock/mock_repository.go

