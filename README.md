# Pike

Generate CRUD gRPC backends from single YAML description.


#### Example

Let's say we want to have backend to keep data on friendly animals.


From following description:

```yaml
name: backend
go_import_path: github.com/sashabaranov/testbackend
entities:
  - name: animal
    fields:
      - {name: id, type: uint32, primary_key: true}
      - {name: name, type: string}
      - {name: age, type: int32}
      - {name: photo_url, type: string}
```

Pike will generate ([example output](https://github.com/sashabaranov/pike/tree/master/examples/testbackend))
* PostgreSQL migration(`CREATE TABLE`) compatible with [migrate](https://github.com/golang-migrate/migrate) tool
* gRPC Protobuf service description with basic Create, Read, Update, Delete operations
* Go implementation of gRPC service 
  * Supports all CRUD SQL queries
  * No additional framework usage. Only depends on `grpc` and `pq`
  * TLS support
  
  
#### Usage

Install pike: `go get github.com/sashabaranov/pike`

Generate all the stuff:


```bash
# Cleanup
PROJ=github.com/sashabaranov/testbackend
DIR=$GOPATH/src/$PROJ
true | rm -rf $DIR

# Generate project
pike examples/animals.yaml


# Generate protobuf
protoc\
	-I $DIR/proto/\
	$DIR/proto/project.proto\
	--go_out=plugins=grpc:$DIR/backend


# Generate certificates
CERT_DIR=$DIR/certs/dev
mkdir -p $CERT_DIR

echo "\n🔖  Generating CA certificate..."
certstrap --depot-path $CERT_DIR init --expires "30 years" --common-name "CA"

echo "\n🔖  Generating server certificate..."
certstrap --depot-path $CERT_DIR request-cert --domain localhost

echo "\n🔖  Signing server certificate with CA..."
certstrap --depot-path $CERT_DIR sign localhost --CA CA
```
