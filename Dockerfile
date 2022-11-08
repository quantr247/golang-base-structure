# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN cd ./cmd/server && go build -o /docker-go-base

EXPOSE 10001

CMD [ "/docker-go-base" ]