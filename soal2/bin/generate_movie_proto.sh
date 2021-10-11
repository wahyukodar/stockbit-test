#!/bin/bash

protoc -I. \
      --proto_path=$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
      --go_out . --go_opt paths=source_relative \
      --go-grpc_out . --go-grpc_opt paths=source_relative \
      --grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
      outbound/movie/proto/service.proto