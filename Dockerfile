FROM golang:1.11.5

WORKDIR /app/gomodtest

COPY go.mod .
COPY go.sum .

RUN go mod download
