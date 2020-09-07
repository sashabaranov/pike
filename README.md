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

Pike is designed to simplify product development. It enables you to give high-level project description at the start and once project is generated Pike does not impose any limitations on what you can do with it. It is not a framework like RoR or Django, just a helpful generator tool.

Pike's name originates from Russian [fairy tale](https://en.wikipedia.org/wiki/At_the_Pike%27s_Behest)

<sub>— По щучьему веленью,
По моему хотенью —
выстройся каменный дворец с золотой крышей…</sub>
