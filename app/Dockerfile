FROM golang:1.22.2-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY .. .

RUN go build -o server main.go

FROM alpine:latest

COPY --from=builder ["/build/server", "/build/.env", "/"]

ENTRYPOINT ["/server"]