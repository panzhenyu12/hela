#!/bin/env bash


all:
	rm -fr bin/*
	PWD=`pwd`
	echo ${GOPATH}
	LASTPATH=${GOPATH}
	GOPATH=${PWD}
	GOPATH="${GOPATH}:${PWD}" gofmt -w src
	GOPATH="${GOPATH}:${PWD}" go install facedeep
	cp src/config.example.json bin/config.json

deps:
	go get "github.com/gorilla/websocket"
	go get "github.com/Shopify/sarama"
	go get "github.com/gorilla/mux"
	go get "github.com/golang/protobuf/proto"
	go get "github.com/satori/go.uuid"
