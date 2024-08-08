package server

import (
	"net"
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

func (s *server) newUserGrpcServer() error {
	lis, err := net.Listen("tcp", s.cfg.GRPC.Port)

	if err != nil {
		return errors.Wrap(err, "net.Listen")
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

	userServiceProto.RegisterUserServiceServer(grpcServer, s.userService)
	grpc_prometheus.Register(grpcServer)

	return grpcServer.Serve(lis)
}
