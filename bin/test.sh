#!/bin/sh


PROJ=github.com/sashabaranov/testbackend

DIR=$GOPATH/src/$PROJ
true | rm -rf $DIR
mkdir -p $DIR/proto
mkdir -p $DIR/backend
mkdir -p $DIR/sql/migrations
mkdir -p $DIR/configs
mkdir -p $DIR/cli

go run cmd/pike.go examples/animals.yaml

protoc\
	-I $DIR/proto/\
	$DIR/proto/project.proto\
	--go_out=plugins=grpc:$DIR/backend


CERT_DIR=$DIR/certs/dev
mkdir -p $CERT_DIR

echo "\n🔖  Generating CA certificate..."
certstrap --depot-path $CERT_DIR init --expires "30 years" --common-name "CA"

echo "\n🔖  Generating server certificate..."
certstrap --depot-path $CERT_DIR request-cert --domain localhost

echo "\n🔖  Signing server certificate with CA..."
certstrap --depot-path $CERT_DIR sign localhost --CA CA