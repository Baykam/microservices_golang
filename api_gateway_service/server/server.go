package server

import (
	"context"
	"os"
	"os/signal"
	"project-microservices/api_gateway_service/config"
	userConnection "project-microservices/api_gateway_service/connections"
	userHttp "project-microservices/api_gateway_service/internal/user/delivery/http"
	u "project-microservices/api_gateway_service/internal/user/service"
	"project-microservices/api_gateway_service/metrics"
	_ "project-microservices/docs"
	"project-microservices/pkg/interceptors"
	productKafka "project-microservices/pkg/kafka"
	"project-microservices/pkg/logger"
	userServiceProto "project-microservices/user_service/proto"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type server struct {
	log         logger.Logger
	cfg         *config.Config
	v           *validator.Validate
	im          interceptors.InterceptorManager
	metrics     metrics.UserMetrics
	userService *u.UserService
	engine      *gin.Engine
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, v: validator.New(), engine: gin.Default()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)
	s.metrics = *metrics.NewUserMetrics(s.cfg)

	userConn, err := userConnection.ConnectGRPCService(ctx, s.cfg.Grpc.UserServicePort, s.im)
	if err != nil {
		return err
	}
	defer userConn.Close()
	userProtoService := userServiceProto.NewUserServiceClient(userConn)
	kafkaProducers := productKafka.NewProducer(s.log, []string{})
	defer kafkaProducers.Close()

	s.userService = u.NewUserService(s.log, *s.cfg, kafkaProducers, userProtoService)

	users := userHttp.NewUserHandlers(*s.cfg, s.log, *s.v, s.engine, &s.metrics, *s.userService)
	users.Run()

	go s.runHttpServer()

	// s.runHealthCheck(ctx)
	// s.runMetrics(cancel)

	<-ctx.Done()
	return nil
}
