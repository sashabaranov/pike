#!/bin/bash
export BACKEND_CONFIG=$GOPATH/src/github.com/sashabaranov/testbackend/configs/dev.yaml
go run $GOPATH/src/github.com/sashabaranov/testbackend/cli/main.go
