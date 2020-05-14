#!/bin/sh

DIR=test_output
true | rm -rf $DIR
mkdir -p $DIR/proto
mkdir -p $DIR/backend
mkdir -p $DIR/sql/migrations

go run cmd/pike.go $DIR

protoc\
	-I $DIR/proto/\
	$DIR/proto/project.proto\
	--go_out=plugins=grpc:$DIR/backend

mkdir -p $DIR/backend