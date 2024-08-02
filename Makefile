.PHONY:

run_api_gateway:
	go run api_gateway_service/cmd/main.go -config=./api_gateway_service/config/config.yaml

run_user_service:
	go run user_service/cmd/main.go -config=./user_service/config/config.yaml

# ==============================================================================
# Go migrate postgresql https://github.com/golang-migrate/migrate

DB_NAME = products
DB_HOST = localhost
DB_PORT = 5432
SSL_MODE = disable

force_db:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations force 1

version_db:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations up

migrate_down:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations down 1


# ==============================================================================
# Proto

proto_kafka_user:
	@echo Generating user messages microservice proto
	cd kafka_protos/user && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. kafka.proto

proto_user:
	@echo Generating user messages microservice proto
	cd user_service/proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. user.proto

# ==============================================================================
# swagger

swagger_user:
	@echo Starting swagger generating
	swag init -g ./api_gateway_service/cmd/main.go -o ./api_gateway_service/docs/user/ --exclude ./api_gateway_service/internal/product

swagger_product:
	@echo Starting swagger generating
	swag init -g ./api_gateway_service/cmd/main.go -o ./api_gateway_service/docs/product/ --exclude ./api_gateway_service/internal/users