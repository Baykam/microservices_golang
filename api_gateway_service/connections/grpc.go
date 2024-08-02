package userConnection

import (
	"context"
	"project-microservices/pkg/interceptors"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	backOffLinear  = 100 * time.Millisecond
	backOffRetries = 3
)

func ConnectGRPCService(ctx context.Context, target string, im interceptors.InterceptorManager) (*grpc.ClientConn, error) {
	opt := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backOffLinear)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
		grpc_retry.WithMax(backOffRetries),
	}
	conn, err := grpc.DialContext(ctx, target,
		grpc.WithUnaryInterceptor(im.ClientRequestLoggerInterceptor()),
		grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opt...)))
	if err != nil {
		return nil, errors.Wrap(err, "grpc.NewClient")
	}
	return conn, nil
}
