APP := jwt-sample
GOPATH ?= $(HOME)/go
#GRPC_GATEWAY_REPO := $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTOC_INC_PATH = $(dir $(shell which protoc))/../include

all: proto/go build

setup:
	GO111MODULE=off go get github.com/golang/protobuf/protoc-gen-go
	#GO111MODULE=off go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

build:
	CGO_ENABLED=0 GOOS=linux go build

proto/go:
	protoc -I. -I$(PROTOC_INC_PATH) \
	--go_out=plugins=grpc:$(GOPATH)/proto/ \
	proto/*.proto

.PHONY: all setup build proto/go