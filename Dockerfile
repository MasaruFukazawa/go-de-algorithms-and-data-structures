FROM golang:1.25

WORKDIR /app

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

#COPY go.mod go.sum ./
