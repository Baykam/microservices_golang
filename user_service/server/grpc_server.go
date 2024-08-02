package server

import (
	"net"
	"project-microservices/user_service/internal/service"
	userServiceProto "project-microservices/user_service/proto"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	maxConnectionIdle = 5 * time.Minute
	gRPCTimeout       = 15 * time.Second
	maxConnectionAge  = 5 * time.Minute
	gRPCTime          = 10 * time.Minute
)

func (s *server) newUserGrpcServer() (func() error, *grpc.Server, error) {
	lis, err := net.Listen("tcp", s.cfg.GRPC.Port)
	if err != nil {
		return nil, nil, errors.Wrap(err, "net.Listen")
	}
	grpcServer := grpc.NewServer(grpc.KeepaliveParams(
		keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle,
			Timeout:           gRPCTimeout,
			MaxConnectionAge:  maxConnectionAge,
			Time:              gRPCTime,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	userService := service.NewUserService(s.postgresRepo, s.cacheRepo)

	userServiceProto.RegisterUserServiceServer(grpcServer, userService)
	grpc_prometheus.Register(grpcServer)

	go func() {
		s.log.Infof("Reader gRPC server is listening on port: %s", s.cfg.GRPC.Port)
		s.log.Fatal(grpcServer.Serve(lis))
	}()
	return lis.Close, grpcServer, nil
}
