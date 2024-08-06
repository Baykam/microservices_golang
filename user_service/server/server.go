package server

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"project-microservices/pkg/logger"
	mongodb "project-microservices/pkg/mongo"
	"project-microservices/pkg/postgres"
	productRedis "project-microservices/pkg/redis"
	"project-microservices/user_service/config"
	"project-microservices/user_service/internal/cache"
	"project-microservices/user_service/internal/repository"
	"project-microservices/user_service/internal/service"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	log          logger.Logger
	cfg          config.Config
	redis        redis.UniversalClient
	sql          *sql.DB
	kafka        *kafka.Conn
	v            *validator.Validate
	mongo        *mongo.Client
	cacheRepo    cache.UserCache
	postgresRepo repository.UserRepository
	userService  service.UserService
}

func NewServer(log logger.Logger, cfg config.Config) *server {
	return &server{log: log, v: validator.New(), cfg: cfg}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.redis = productRedis.ConnRedis(s.cfg.Redis)
	mongoClient, err := mongodb.NewMongoDbConn(ctx, s.cfg.Mongo)
	if err != nil {
		return errors.Wrap(err, "mongodb.NewMongoDbConn")
	}
	s.mongo = mongoClient
	sql, err := postgres.ConnPostgres(ctx, s.cfg.Postgres)
	if err != nil {
		return err
	}
	s.sql = sql

	s.cacheRepo = cache.NewUserCache(s.log, &s.cfg, s.redis)
	s.postgresRepo = repository.NewUserRepository(s.sql)

	s.userService = service.NewUserService(s.postgresRepo, s.cacheRepo, s.cfg)

	// userMessageProcessor := userKafkaConn.NewUserMessagesProcessor(s.log, &s.cfg, s.v, s.userService)

	// kafkaConsumerGroup := productKafka.NewConsumerGroup(s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupId, s.log)
	// go kafkaConsumerGroup.ConsumeTopic(ctx, s.getConsumerGroupTopics(), userKafkaConn.PoolSize, userMessageProcessor.ProcessMessages)

	// if err := s.connectKafkaBrokers(ctx); err != nil {
	// 	return errors.Wrap(err, "server.connectKafkaBrokers")
	// }
	// defer s.kafka.Close()

	// closeGrpcServer, grpcServer, err = s.newUserGrpcServer()
	// if err != nil {
	// 	return err
	// }
	// defer closeGrpcServer()
	// defer grpcServer.GracefulStop()

	go func() {
		if err := s.newUserGrpcServer(); err != nil {
			s.log.Fatal(err)
		}
	}()

	<-ctx.Done()
	s.log.Info("Shutting down gRPC server...")
	return nil
}
