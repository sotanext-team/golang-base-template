## Setup

### Install dependencies

- GoLang: version 1.17
- Docker: https://docs.docker.com/get-docker/
- Docker Compose: https://docs.docker.com/compose/install/
- After installing docker, create a network for es-hs by using the command

```bash
docker network create es-hs
```

## Project structure

Referring from these repositories

- https://github.com/bxcodec/go-clean-arch
- https://github.com/golang-standards/project-layout

## Running

### Run locally

```bash
docker-compose -f deployments/docker-compose.yml -p app-api --env-file .env up -d
```

```bash
# Run the postgres service
docker run -d \
    --name postgres_local \
    -e POSTGRES_PASSWORD=secret \
    -p 54321:5432 \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /pgdata:/var/lib/postgresql/data \
    postgres:12-alpine
make watch
```

Create database in docker container

```bash
docker exec -it postgres_local psql -U postgres -c "CREATE DATABASE app_api"
```

If you update Dockerfile, rebuild by the command

```bash
make down
make watch
```

### Seeding data

```bash
make seeding
```

## Folder structure

```
|__ env.example.yml
|__ env.yml
|__ server.go
|__ bin
|__ helpers
|__ models
    |__ database.go
    |__ product.go
    |__ product_variant.go
|__ modules
    |__ product
        |__ grpc
        |__ http
        |__ graphql
            |__ product_graphql.go
            |__ product_type.go
        |__ repository
            |__ product_repository.go
            |__ product_repository_test.go
        |__ usecase
            |__ product_usercase.go
            |__ product_usercase_test.go
        |__ repository.go
        |__ usecase.go
    |__ product_variant
        |__ graphql
            |__ product_variant_graphql.go
            |__ product_variant_type.go
        |__ repository
            |__ product_variant_repository.go
        |__ usecase
            |__ product_variant_usercase.go
        |__ repository.go
        |__ usecase.go
|__ public
    |__ graphql
```

## Using ent and gqlgen

### Create new ent model

```bash
ent init Car
```
