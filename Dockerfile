FROM docker.io/library/golang:1.24.4 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build ./cmd/api

FROM docker.io/library/alpine:latest

WORKDIR /root/

COPY --from=builder /src .

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./api"]
