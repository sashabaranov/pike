package backend

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"runtime/debug"
	// "github.com/getsentry/sentry-go"
)

func panicRecover(p interface{}) error {
	reportError(fmt.Sprintf("Panic happened: %s", p), nil)
	log.Printf("Panic ouccred: %s", p)
	log.Print("Stacktrace from panic: \n" + string(debug.Stack()))
	return status.Errorf(codes.Internal, "Internal error")
}

func RunServer() {
	config := LoadConfig()
	server := NewServerFromConfig(config)

	// if config.SentryDSN != "" {
	// 	sentry.Init(sentry.ClientOptions{Dsn: config.SentryDSN})
	// }

	grpcServer := grpc.NewServer(makeServerOptions(config)...)
	RegisterBackendServer(grpcServer, server)

	lis, err := net.Listen("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server started on %s", config.ListenAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func makeServerOptions(config ServerConfig) []grpc.ServerOption {
	transportCredentials, err := getTransportCredentials(config)
	if err != nil {
		log.Fatalf("failed to get credentials: %v", err)
	}

	logger := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	logrusEntry := logrus.NewEntry(logger)

	recoverOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(panicRecover),
	}
	unaryServerInterceptors := []grpc.UnaryServerInterceptor{
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_logrus.UnaryServerInterceptor(logrusEntry),
		grpc_auth.UnaryServerInterceptor(nil),
		grpc_recovery.UnaryServerInterceptor(recoverOptions...),
	}

	return []grpc.ServerOption{
		grpc.MaxRecvMsgSize(config.MaxMessageSizeBytes),
		grpc.MaxSendMsgSize(config.MaxMessageSizeBytes),
		grpc.Creds(*transportCredentials),
		grpc_middleware.WithUnaryServerChain(unaryServerInterceptors...),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
			grpc_auth.StreamServerInterceptor(nil),
			grpc_recovery.StreamServerInterceptor(recoverOptions...),
		),
	}
}