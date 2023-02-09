.PHONY: run
run:
	swag init -g cmd/main.go && docker-compose up && goose-i
migrate-status:
	docker exec mcrsrv-book $(MAKE) migration-status -C migration
migrate-up:
	docker exec mcrsrv-book $(MAKE) migration-up -C migration
migrate-down:
	docker exec mcrsrv-book $(MAKE) migration-down -C migration
migrate-create:
	docker exec mcrsrv-book $(MAKE) migration-create -C migration
documentation-create:
	swag init -g cmd/main.go
test:
	go test ./tests/...
goose-i:
	curl -fsSL \
        https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
        sh
refresh-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 		internal/controllers/findBook/gRPC/findBook.proto && \
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	 	internal/controllers/listBook/gRPC/listBook.proto && \
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/controllers/createBook/gRPC/createBook.proto && \
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/controllers/updateBook/gRPC/updateBook.proto && \
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/controllers/deleteBook/gRPC/deleteBook.proto

