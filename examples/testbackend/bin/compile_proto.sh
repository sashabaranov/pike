#!/bin/bash
protoc \
    --go_out=plugins=grpc:$GOPATH/src/github.com/sashabaranov/testbackend/backend \
    -I $GOPATH/src \
    -I $GOPATH/src/github.com/sashabaranov/testbackend/proto \
    $GOPATH/src/github.com/sashabaranov/testbackend/proto/backend.proto
