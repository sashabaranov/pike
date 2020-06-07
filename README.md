# Pike

Generate CRUD gRPC backends from single YAML description.


Pike generates:

* gRPC Protobuf service description with basic Create, Read, Update, Delete operations
* Go implementation of gRPC service 
  * Supports all CRUD SQL queries
  * No additional framework usage. Only depends on `grpc` and `pq`
  * TLS support
* PostgreSQL migration(`CREATE TABLE`) compatible with [migrate](https://github.com/golang-migrate/migrate) tool

#### Usage

Install: `go get github.com/sashabaranov/pike`

Run: `pike project.yaml`

#### Example

![](https://i.imgur.com/DVgPfu8.png)

Generated project can be found ([here](https://github.com/sashabaranov/pike/tree/master/examples/testbackend))
