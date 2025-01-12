FROM golang:latest AS builder
RUN apt update && apt install git -y

RUN mkdir $GOPATH/src/server
ADD ./prometheus.go $GOPATH/src/server

WORKDIR $GOPATH/src/server
RUN go mod init
RUN go mod tidy
RUN go mod download
RUN mkdir /pro
RUN go build -o /pro/server prometheus.go

FROM debian:stable-slim

RUN mkdir -p /pro
COPY --from=builder /pro/server /pro/server
EXPOSE 1234
WORKDIR /pro
CMD [ "/pro/server" ]