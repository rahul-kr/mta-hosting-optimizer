FROM golang:alpine

WORKDIR /app

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./app