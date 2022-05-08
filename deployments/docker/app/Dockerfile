FROM golang:1.18-alpine

RUN apk add bash ca-certificates git gcc g++ libc-dev curl

ENV CGO_ENABLED=0

ARG GO111MODULE=on

CMD ["./scripts/app.sh"]
