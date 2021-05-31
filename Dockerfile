FROM golang:latest

RUN mkdir /go/src/money-app

ENV GOPATH /root/.go

WORKDIR /go/src/money-app