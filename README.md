## Setup

### Install dependencies

-   GoLang: version 1.18
-   Docker: https://docs.docker.com/get-docker/
-   Docker Compose: https://docs.docker.com/compose/install/

## Project structure

Referring from these repositories

-   https://github.com/bxcodec/go-clean-arch
-   https://github.com/golang-standards/project-layout

## Running

### Run locally

```bash
# Run the postgres service
docker run -d \
    --name postgres-local \
    -e POSTGRES_PASSWORD=secret \
    -p 54321:5432 \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /pgdata:/var/lib/postgresql/data \
    postgres:12-alpine
source .env
go run server.go
```

Create database in docker container

```bash
docker exec -it postgres-local psql -U postgres -c "CREATE DATABASE app_api"
```

or running the command:

```bash
make initdb
```
