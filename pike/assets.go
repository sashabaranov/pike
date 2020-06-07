package pike

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsfd0d21aba52c2a14bfda27f857b21eb4779c63d1 = "#!/bin/bash\nprotoc \\\n    --go_out=plugins=grpc:$GOPATH/src/{{.GoImportPath}}/{{.Name}} \\\n    -I $GOPATH/src \\\n    -I $GOPATH/src/{{.GoImportPath}}/proto \\\n    $GOPATH/src/{{.GoImportPath}}/proto/{{.Name}}.proto\n"
var _Assetsb8cf25a6ea3c69d230dc23047c387a2155c41022 = "package {{.Name}}\n\nimport (\n\t\"database/sql\"\n\t_ \"github.com/lib/pq\"\n\t\"time\"\n)\n\ntype PostgreStorage struct {\n\tdb  *sql.DB\n\turi string\n}\n\nfunc NewPostgreStorage(uri string) (*PostgreStorage, error) {\n\tret := &PostgreStorage{\n\t\turi: uri,\n\t}\n\n\terr := ret.Connect()\n\treturn ret, err\n}\n\nfunc (storage *PostgreStorage) Connect() error {\n\tdb, err := sql.Open(\"postgres\", storage.uri)\n\tif err != nil {\n\t\treturn err\n\t}\n\n\tstorage.db = db\n\tdb.SetMaxOpenConns(3)\n\tdb.SetMaxIdleConns(3)\n\tdb.SetConnMaxLifetime(30 * time.Minute)\n\treturn err\n}\n\n"
var _Assets496fffaa628b669f1f90a970819f0d8b15139026 = "package {{.Name}}\n\nimport (\n\t\"gopkg.in/yaml.v2\"\n\t\"log\"\n\t\"os\"\n\t\"io/ioutil\"\n)\n\ntype Server struct {\n\tUnimplemented{{.ProtoCapsName}}Server\n\tstorage *PostgreStorage\n}\n\ntype ServerConfig struct {\n\tListenAddr string `yaml:\"listen_on\"`\n\tDatabaseURI string `yaml:\"db_uri\"`\n\n\tMaxMessageSizeBytes int `yaml:\"max_message_size_bytes\"`\n\n\t// SentryDSN string `yaml:\"sentry_dsn\"`\n\n\tServerCert string `yaml:\"server_cert\"`\n\tServerKey  string `yaml:\"server_key\"`\n\tCACert     string `yaml:\"ca_cert\"`\n}\n\nfunc LoadConfig() ServerConfig {\n\tconfigPath := os.Getenv(\"{{.ConfigEnvVariable}}\")\n\tcontent, err := ioutil.ReadFile(configPath)\n\tif err != nil {\n\t\tlog.Fatalf(\n\t\t\t\"Error loading config from {{.ConfigEnvVariable}}=%s: %v\",\n\t\t\tconfigPath,\n\t\t\terr,\n\t\t)\n\t}\n\n\tconfig := ServerConfig{}\n\terr = yaml.Unmarshal([]byte(content), &config)\n\tif err != nil {\n\t\tlog.Fatalf(\"Error parsing config: %v\", err)\n\t}\n\treturn config\n}\n\nfunc NewServerFromConfig(cfg ServerConfig) *Server {\n\tstorage, err := NewPostgreStorage(cfg.DatabaseURI)\n\tif err != nil {\n\t\tlog.Fatalf(\"Could not create storage: %v\", err)\n\t}\n\n\treturn &Server{\n\t\tstorage: storage,\n\t}\n}\n\n\nfunc (s *Server) Cleanup() {\n\ts.storage.db.Close()\n}"
var _Assets336cb40f54e146e1a5477fc39dc37464e5a4f48e = "package {{.Name}}\n\nimport (\n\t\"context\"\n)\n\nfunc (s *Server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {\n\t// allowed_endpoints := map[string]bool{\n\t// \t\"/backend.Backend/CreateUser\":   true,\n\t// }\n\n\t// if allow, ok := allowed_endpoints[fullMethodName]; allow && ok {\n\t// \treturn ctx, nil\n\t// }\n\treturn ctx, nil\n\n}\n"
var _Assetsb075389a6a17cbec666740f4aa4446e7b5aa84b9 = "package {{.Name}}\n\nimport (\n\t\"context\"\n\t\"google.golang.org/grpc/codes\"\n\t\"google.golang.org/grpc/status\"\n)\n\n{{- range  .Entities }}\n\nfunc (s *Server) Create{{.ProtoCapsName}}(ctx context.Context, req *Create{{.ProtoCapsName}}Request) (*Create{{.ProtoCapsName}}Response, error) {\n\tcreated, err := s.storage.Create{{.ProtoCapsName}}(req.{{.ProtoCapsName}})\n\tif err != nil {\n\t\treportError(\"Error in Create{{.ProtoCapsName}}\", err)\n\t\treturn nil, status.Error(codes.Internal, \"Internal error\")\n\t}\n\treturn &Create{{.ProtoCapsName}}Response{\n\t\tCreated: created,\n\t}, nil\n}\n\nfunc (s *Server) Get{{.ProtoCapsName}}(ctx context.Context, req *Get{{.ProtoCapsName}}Request) (*Get{{.ProtoCapsName}}Response, error) {\n\tret, err := s.storage.Get{{.ProtoCapsName}}(req.{{.PrimaryKeyField.GoName}})\n\tif err != nil {\n\t\treportError(\"Error in Get{{.ProtoCapsName}}\", err)\n\t\treturn nil, status.Error(codes.Internal, \"Internal error\")\n\t}\n\n\treturn &Get{{.ProtoCapsName}}Response{\n\t\t{{.ProtoCapsName}}: ret,\n\t}, nil\n}\n\nfunc (s *Server) Update{{.ProtoCapsName}}(ctx context.Context, req *Update{{.ProtoCapsName}}Request) (*Update{{.ProtoCapsName}}Response, error) {\n\tupdated, err := s.storage.Update{{.ProtoCapsName}}(req.Updated)\n\tif err != nil {\n\t\treportError(\"Error in Update{{.ProtoCapsName}}\", err)\n\t\treturn nil, status.Error(codes.Internal, \"Internal error\")\n\t}\n\n\treturn &Update{{.ProtoCapsName}}Response{\n\t\tResult: updated,\n\t}, nil\n}\n\nfunc (s *Server) Delete{{.ProtoCapsName}}(ctx context.Context, req *Delete{{.ProtoCapsName}}Request) (*Delete{{.ProtoCapsName}}Response, error) {\n\terr := s.storage.Delete{{.ProtoCapsName}}(req.{{.PrimaryKeyField.GoName}})\n\tif err != nil {\n\t\treportError(\"Error in Delete{{.ProtoCapsName}}\", err)\n\t\treturn nil, status.Error(codes.Internal, \"Internal error\")\n\t}\n\treturn &Delete{{.ProtoCapsName}}Response{}, nil\n}\n\n\n{{- end }}\n"
var _Assetsd3ee4c59e20ce424f3740717d241532718187045 = "{{- range .Entities }}DROP TABLE IF EXISTS {{.SQLTableName}};\n{{end}}"
var _Assetsfadd1e5514b9a8efffac4384eda09894fff76c2f = "#!/bin/bash\nexport {{.ConfigEnvVariable}}=$GOPATH/src/{{.GoImportPath}}/configs/dev.yaml\ngo run $GOPATH/src/{{.GoImportPath}}/cli/main.go\n"
var _Assets3ce337a55365306ec3e0213b1e478eb0461183c3 = "package {{.Name}}\n"
var _Assets9eea5c7000f3012e7b4f51ff373f2898a805bfb7 = "package {{.Name}}\n\nimport (\n\t\"crypto/tls\"\n\t\"crypto/x509\"\n\t\"fmt\"\n\t\"github.com/grpc-ecosystem/go-grpc-middleware\"\n\t\"github.com/grpc-ecosystem/go-grpc-middleware/auth\"\n\t\"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus\"\n\t\"github.com/grpc-ecosystem/go-grpc-middleware/recovery\"\n\t\"github.com/grpc-ecosystem/go-grpc-middleware/tags\"\n\t\"github.com/sirupsen/logrus\"\n\t\"google.golang.org/grpc\"\n\t\"google.golang.org/grpc/codes\"\n\t\"google.golang.org/grpc/credentials\"\n\t\"google.golang.org/grpc/status\"\n\t\"io/ioutil\"\n\t\"log\"\n\t\"net\"\n\t\"os\"\n\t\"runtime/debug\"\n\t// \"github.com/getsentry/sentry-go\"\n)\n\nfunc panicRecover(p interface{}) error {\n\treportError(fmt.Sprintf(\"Panic happened: %s\", p), nil)\n\tlog.Printf(\"Panic ouccred: %s\", p)\n\tlog.Print(\"Stacktrace from panic: \\n\" + string(debug.Stack()))\n\treturn status.Errorf(codes.Internal, \"Internal error\")\n}\n\nfunc RunServer() {\n\tconfig := LoadConfig()\n\tserver := NewServerFromConfig(config)\n\tdefer server.Cleanup()\n\n\t// if config.SentryDSN != \"\" {\n\t// \tsentry.Init(sentry.ClientOptions{Dsn: config.SentryDSN})\n\t// }\n\n\tgrpcServer := grpc.NewServer(makeServerOptions(config)...)\n\tRegister{{.ProtoCapsName}}Server(grpcServer, server)\n\n\tlis, err := net.Listen(\"tcp\", config.ListenAddr)\n\tif err != nil {\n\t\tlog.Fatalf(\"failed to listen: %v\", err)\n\t}\n\n\tlog.Printf(\"Server started on %s\", config.ListenAddr)\n\tif err := grpcServer.Serve(lis); err != nil {\n\t\tlog.Fatalf(\"failed to serve: %v\", err)\n\t}\n}\n\nfunc makeServerOptions(config ServerConfig) []grpc.ServerOption {\n\ttransportCredentials, err := getTransportCredentials(config)\n\tif err != nil {\n\t\tlog.Fatalf(\"failed to get credentials: %v\", err)\n\t}\n\n\tlogger := &logrus.Logger{\n\t\tOut:       os.Stdout,\n\t\tFormatter: new(logrus.TextFormatter),\n\t\tHooks:     make(logrus.LevelHooks),\n\t\tLevel:     logrus.DebugLevel,\n\t}\n\tlogrusEntry := logrus.NewEntry(logger)\n\n\trecoverOptions := []grpc_recovery.Option{\n\t\tgrpc_recovery.WithRecoveryHandler(panicRecover),\n\t}\n\tunaryServerInterceptors := []grpc.UnaryServerInterceptor{\n\t\tgrpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),\n\t\tgrpc_logrus.UnaryServerInterceptor(logrusEntry),\n\t\tgrpc_auth.UnaryServerInterceptor(nil),\n\t\tgrpc_recovery.UnaryServerInterceptor(recoverOptions...),\n\t}\n\n\treturn []grpc.ServerOption{\n\t\tgrpc.MaxRecvMsgSize(config.MaxMessageSizeBytes),\n\t\tgrpc.MaxSendMsgSize(config.MaxMessageSizeBytes),\n\t\tgrpc.Creds(*transportCredentials),\n\t\tgrpc_middleware.WithUnaryServerChain(unaryServerInterceptors...),\n\t\tgrpc_middleware.WithStreamServerChain(\n\t\t\tgrpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),\n\t\t\tgrpc_logrus.StreamServerInterceptor(logrusEntry),\n\t\t\tgrpc_auth.StreamServerInterceptor(nil),\n\t\t\tgrpc_recovery.StreamServerInterceptor(recoverOptions...),\n\t\t),\n\t}\n}\n\nfunc getTransportCredentials(cfg ServerConfig) (*credentials.TransportCredentials, error) {\n\tpeerCert, err := tls.LoadX509KeyPair(cfg.ServerCert, cfg.ServerKey)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tcaCert, err := ioutil.ReadFile(cfg.CACert)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tcaCertPool := x509.NewCertPool()\n\tcaCertPool.AppendCertsFromPEM(caCert)\n\ttc := credentials.NewTLS(&tls.Config{\n\t\tCertificates: []tls.Certificate{peerCert},\n\t\tClientCAs:    caCertPool,\n\t})\n\n\treturn &tc, nil\n}\n"
var _Assets309897f3245c4c736c4e5707577a4c66105fff4b = "listen_on: 0.0.0.0:8000\ndb_uri: postgres://localhost:5432/{{.Name}}?sslmode=disable\nmax_message_size_bytes: 500000000\nserver_cert: certs/dev/localhost.crt\nserver_key: certs/dev/localhost.key\nca_cert: certs/dev/CA.crt"
var _Assetsdbfac68590dd233d9aa4715578dcaf02f8c7bb6d = "syntax = \"proto3\";\n\npackage {{.Name}};\noption go_package = \".;{{.Name}}\";\n\nservice {{.ProtoCapsName}} {\n    {{- range .Entities }}\n    rpc Create{{.ProtoCapsName}}(Create{{.ProtoCapsName}}Request) returns (Create{{.ProtoCapsName}}Response);\n    rpc Get{{.ProtoCapsName}}(Get{{.ProtoCapsName}}Request) returns (Get{{.ProtoCapsName}}Response);\n    rpc Update{{.ProtoCapsName}}(Update{{.ProtoCapsName}}Request) returns (Update{{.ProtoCapsName}}Response);\n    rpc Delete{{.ProtoCapsName}}(Delete{{.ProtoCapsName}}Request) returns (Delete{{.ProtoCapsName}}Response);\n\n    {{- end }}\n}\n\n{{- range .Entities }}\n\nmessage {{.ProtoCapsName}} {\n    {{- range $index, $field := .Fields}}\n    {{$field.Type}} {{$field.Name}} = {{inc $index}};\n    {{- end}}\n}\n\nmessage Create{{.ProtoCapsName}}Request {\n    {{.ProtoCapsName}} {{.Name}} = 1;\n}\n\nmessage Create{{.ProtoCapsName}}Response {\n    {{.ProtoCapsName}} created = 1;\n}\n\nmessage Get{{.ProtoCapsName}}Request {\n    {{.PrimaryKeyField.Type}} {{.PrimaryKeyField.Name}} = 1;\n}\n\nmessage Get{{.ProtoCapsName}}Response {\n    {{.ProtoCapsName}} {{.Name}} = 1;\n}\n\nmessage Update{{.ProtoCapsName}}Request {\n    {{.ProtoCapsName}} updated = 1;\n}\n\nmessage Update{{.ProtoCapsName}}Response {\n    {{.ProtoCapsName}} result = 1;\n}\n\nmessage Delete{{.ProtoCapsName}}Request {\n    {{.PrimaryKeyField.Type}} {{.PrimaryKeyField.Name}} = 1;\n}\n\nmessage Delete{{.ProtoCapsName}}Response {\n\n}\n{{- end }}"
var _Assets16bc4c6adbccea364db29618459c21c6189fd764 = "package {{.Name}}\n\nimport (\n\t// \"fmt\"\n\t// \"github.com/getsentry/sentry-go\"\n\t\"log\"\n)\n\nfunc reportError(message string, err error) {\n\t// sentry.CaptureException(fmt.Errorf(\"%s: %v\", message, err))\n\tlog.Printf(\"🚨 %s – %v\", message, err)\n}\n"
var _Assetscd19369b6b09fea86477e4d954b8654ba2496892 = "package main\n\nimport (\n\t\"{{.GoImportPath}}/{{.Name}}\"\n)\n\nfunc main() {\n\t{{.Name}}.RunServer()\n}\n"
var _Assets6680918d504a6d0f79c44f574a7412385689de3f = "{{- range $entityIx, $entity := .Entities }}CREATE TABLE {{$entity.SQLTableName}} (\n{{- range $ix, $field := $entity.Fields}}\n\t{{$field.Name}} {{$field.SQLType}},\n{{- end }}\n\tcreated timestamp(0) without time zone default (now() at time zone 'utc')\n);\n\n{{end}}"
var _Assets26aa94d6419f9f56cb08934ef33c03bb4141abbf = "package {{.Name}}\n\nimport ()\n\n{{- range $entityIx, $entity := .Entities }}\n\nfunc (storage *PostgreStorage) Create{{$entity.ProtoCapsName}}(in *{{$entity.ProtoCapsName}}) (*{{$entity.ProtoCapsName}}, error) {\n\tstmt, err := storage.db.Prepare(`\n\t\tINSERT INTO {{$entity.SQLTableName}}(\n\t\t\t{{- range $ix, $field := $entity.NonPrimaryKeyFields}}\n\t\t\t{{$field.Name}}{{- if not (last $ix $entity.NonPrimaryKeyFields)}},{{- end}}\n\t\t\t{{- end }}\n\t\t)\n\t\tVALUES ({{- range $ix, $field := $entity.NonPrimaryKeyFields}}${{inc $ix}}{{- if not (last $ix $entity.NonPrimaryKeyFields)}},{{- end}}{{- end }})\n\t\tRETURNING {{$entity.PrimaryKeyField.Name}};\n\t`)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer stmt.Close()\n\n\terr = stmt.QueryRow(\n\t\t{{- range $ix, $field := $entity.NonPrimaryKeyFields}}\n\t\t in.{{$field.GoName}},\n\t\t{{- end }}\n\t).Scan(\n\t\t&(in.{{$entity.PrimaryKeyField.GoName}}),\n\t)\n\n\treturn in, err\n}\n\n\nfunc (storage *PostgreStorage) Delete{{$entity.ProtoCapsName}}(id {{$entity.PrimaryKeyField.GoType}}) error {\n\tstmt, err := storage.db.Prepare(\"DELETE FROM {{$entity.SQLTableName}} WHERE {{$entity.PrimaryKeyField.Name}}=$1;\")\n\tif err != nil {\n\t\treturn err\n\t}\n\tdefer stmt.Close()\n\n\t_, err = stmt.Exec(id)\n\tif err != nil {\n\t\treturn err\n\t}\n\n\treturn nil\n}\n\n\nfunc (storage *PostgreStorage) Get{{$entity.ProtoCapsName}}(id {{$entity.PrimaryKeyField.GoType}}) (*{{$entity.ProtoCapsName}}, error) {\n\tstmt, err := storage.db.Prepare(`\n\t\tSELECT\n\t\t\t{{- range $ix, $field := $entity.Fields}}\n\t\t\t{{$field.Name}}{{- if not (last $ix $entity.Fields)}},{{- end}}\n\t\t\t{{- end }}\n\t\tFROM {{$entity.SQLTableName}}\n\t\tWHERE {{$entity.PrimaryKeyField.Name}}=$1;\n\t`)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer stmt.Close()\n\n\tret := &{{$entity.ProtoCapsName}}{}\n\terr = stmt.QueryRow(id).Scan(\n\t\t{{- range $ix, $field := $entity.Fields}}\n\t\t&ret.{{$field.GoName}},\n\t\t{{- end }}\n\t)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\treturn ret, nil\n}\n\n\nfunc (storage *PostgreStorage) Update{{$entity.ProtoCapsName}}(updated *{{$entity.ProtoCapsName}}) (*{{$entity.ProtoCapsName}}, error) {\n\ttx, err := storage.db.Begin()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tstmt, err := tx.Prepare(`\n\t\tUPDATE {{$entity.SQLTableName}}\n\t\tSET\n\t\t\t{{- range $ix, $field := $entity.NonPrimaryKeyFields}}\n\t\t\t{{$field.Name}}=${{inc $ix}}{{- if not (last $ix $entity.NonPrimaryKeyFields)}},{{- end}}\n\t\t\t{{- end }}\n\t\tWHERE\n\t\t\t{{$entity.PrimaryKeyField.Name}}=${{inc (len $entity.Fields)}}\n\t\tRETURNING\n\t\t\t{{- range $ix, $field := $entity.Fields}}\n\t\t\t{{$field.Name}}{{- if not (last $ix $entity.Fields)}},{{- end}}\n\t\t\t{{- end }}\n\t\t;\n\t`)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer stmt.Close()\n\n\tret := &{{$entity.ProtoCapsName}}{}\n\terr = stmt.QueryRow(\n\t\t{{- range $ix, $field := $entity.NonPrimaryKeyFields}}\n\t\tupdated.{{$field.GoName}},\n\t\t{{- end }}\n\t).Scan(\n\t\t{{- range $ix, $field := $entity.Fields}}\n\t\t&ret.{{$field.GoName}},\n\t\t{{- end }}\n\t)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\terr = tx.Commit()\n\tif err != nil {\n\t\ttx.Rollback()\n\t}\n\n\treturn ret, err\n}\n{{- end }}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"config.go.tmplt", "compile_proto.sh.tmplt", "storage.go.tmplt", "initial_migration.down.sql.tmplt", "server.go.tmplt", "auth.go.tmplt", "run.go.tmplt", "config.yaml.tmplt", "project_proto.tmplt", "launcher.go.tmplt", "report_error.go.tmplt", "server_entity.go.tmplt", "initial_migration.up.sql.tmplt", "run.sh.tmplt", "storage_entity.go.tmplt"}}, map[string]*assets.File{
	"/templates/run.go.tmplt": &assets.File{
		Path:     "/templates/run.go.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503711632307),
		Data:     []byte(_Assets9eea5c7000f3012e7b4f51ff373f2898a805bfb7),
	}, "/templates/config.yaml.tmplt": &assets.File{
		Path:     "/templates/config.yaml.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503710239969),
		Data:     []byte(_Assets309897f3245c4c736c4e5707577a4c66105fff4b),
	}, "/templates/project_proto.tmplt": &assets.File{
		Path:     "/templates/project_proto.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503711147916),
		Data:     []byte(_Assetsdbfac68590dd233d9aa4715578dcaf02f8c7bb6d),
	}, "/templates/report_error.go.tmplt": &assets.File{
		Path:     "/templates/report_error.go.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503711361487),
		Data:     []byte(_Assets16bc4c6adbccea364db29618459c21c6189fd764),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1591515737, 1591515737022760230),
		Data:     nil,
	}, "/templates/config.go.tmplt": &assets.File{
		Path:     "/templates/config.go.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503710035238),
		Data:     []byte(_Assets3ce337a55365306ec3e0213b1e478eb0461183c3),
	}, "/templates/initial_migration.up.sql.tmplt": &assets.File{
		Path:     "/templates/initial_migration.up.sql.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503710647804),
		Data:     []byte(_Assets6680918d504a6d0f79c44f574a7412385689de3f),
	}, "/templates/storage_entity.go.tmplt": &assets.File{
		Path:     "/templates/storage_entity.go.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503712927256),
		Data:     []byte(_Assets26aa94d6419f9f56cb08934ef33c03bb4141abbf),
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1591450503, 1591450503712839189),
		Data:     nil,
	}, "/templates/launcher.go.tmplt": &assets.File{
		Path:     "/templates/launcher.go.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503710879856),
		Data:     []byte(_Assetscd19369b6b09fea86477e4d954b8654ba2496892),
	}, "/templates/server.go.tmplt": &assets.File{
		Path:     "/templates/server.go.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503712183960),
		Data:     []byte(_Assets496fffaa628b669f1f90a970819f0d8b15139026),
	}, "/templates/auth.go.tmplt": &assets.File{
		Path:     "/templates/auth.go.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591518388, 1591518388956714616),
		Data:     []byte(_Assets336cb40f54e146e1a5477fc39dc37464e5a4f48e),
	}, "/templates/server_entity.go.tmplt": &assets.File{
		Path:     "/templates/server_entity.go.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503712436134),
		Data:     []byte(_Assetsb075389a6a17cbec666740f4aa4446e7b5aa84b9),
	}, "/templates/compile_proto.sh.tmplt": &assets.File{
		Path:     "/templates/compile_proto.sh.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503709843489),
		Data:     []byte(_Assetsfd0d21aba52c2a14bfda27f857b21eb4779c63d1),
	}, "/templates/storage.go.tmplt": &assets.File{
		Path:     "/templates/storage.go.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503712695202),
		Data:     []byte(_Assetsb8cf25a6ea3c69d230dc23047c387a2155c41022),
	}, "/templates/initial_migration.down.sql.tmplt": &assets.File{
		Path:     "/templates/initial_migration.down.sql.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503710420645),
		Data:     []byte(_Assetsd3ee4c59e20ce424f3740717d241532718187045),
	}, "/templates/run.sh.tmplt": &assets.File{
		Path:     "/templates/run.sh.tmplt",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1591450503, 1591450503711915281),
		Data:     []byte(_Assetsfadd1e5514b9a8efffac4384eda09894fff76c2f),
	}}, "")
