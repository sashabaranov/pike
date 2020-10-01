# Pike

Generate CRUD gRPC backends from single YAML description.

Check out **[Playground](https://backend-playground.transcendent.app/)**!

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

<img src="https://i.imgur.com/DVgPfu8.png" width="70%" />

Generated project can be found [here](https://github.com/sashabaranov/pike/tree/master/examples)


#### Philosophy


> A complex system that works is invariably found to have evolved from a simple system that worked. A complex system designed from scratch never works and cannot be patched up to make it work. You have to start over with a working simple system. 
>
> — Gall's Law

Pike let's you create simple systems quickly and does not impose any limitations afterwards. 

Pike's name originates from Russian [fairy tale](https://en.wikipedia.org/wiki/At_the_Pike%27s_Behest)

<sub>— По щучьему веленью,
По моему хотенью —
выстройся каменный дворец с золотой крышей…</sub>
