#!/bin/sh


PROJ=github.com/sashabaranov/testbackend
DIR=$GOPATH/src/$PROJ
true | rm -rf $DIR

go run pike.go examples/animals.yaml

CERT_DIR=$DIR/certs/dev
mkdir -p $CERT_DIR

echo "\n🔖  Generating CA certificate..."
certstrap --depot-path $CERT_DIR init --expires "30 years" --common-name "CA"

echo "\n🔖  Generating server certificate..."
certstrap --depot-path $CERT_DIR request-cert --domain localhost

echo "\n🔖  Signing server certificate with CA..."
certstrap --depot-path $CERT_DIR sign localhost --CA CA
