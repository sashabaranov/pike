# Pike

Generate CRUD gRPC backends from single YAML description.


#### Usage

Install: `go get github.com/sashabaranov/pike`

Run: `pike project.yaml`

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
      - {name: name, type: string, sql_type: "VARCHAR(128)"}
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
  
Pike's output:

![](https://i.imgur.com/k7htnKq.png)
