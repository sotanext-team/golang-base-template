app:
	go run server.go -task server

grpc:
	go run server.go -task service

watch:
	./scripts/app.sh

watch-service:
	./scripts/service.sh

build:
	GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"

test:
	godotenv go test ./...

seeding:
	go run cmd/seeding/main.go

initdb:
	docker exec -it postgres_local psql -U postgres -c "CREATE DATABASE app_api"

generate:
	go generate ./...

generate-ent:
	go generate ./ent

generate-gql:
	go generate ./graph
# build:
# 	# TODO
