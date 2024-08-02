package config

import (
	"fmt"
	"os"
	"project-microservices/pkg/constants"
	productKafka "project-microservices/pkg/kafka"
	"project-microservices/pkg/logger"
	mongodb "project-microservices/pkg/mongo"
	"project-microservices/pkg/postgres"
	"project-microservices/pkg/probes"
	productRedis "project-microservices/pkg/redis"
	"project-microservices/pkg/tracing"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName      string               `mapstructure:"serviceName"`
	Logger           *logger.Config       `mapstructure:"logger"`
	GRPC             Grpc                 `mapstructure:"grpc"`
	Postgres         *postgres.Config     `mapstructure:"postgres"`
	Mongo            *mongodb.Config      `mapstructure:"mongo"`
	Redis            *productRedis.Config `mapstructure:"redis"`
	Kafka            *productKafka.Config `mapstructure:"kafka"`
	MongoCollections MongoCollections     `mapstructure:"mongoCollections"`
	Probes           probes.Config        `mapstructure:"probes"`
	ServiceSettings  ServiceSettings      `mapstructure:"serviceSettings"`
	Jaeger           *tracing.Config      `mapstructure:"jaeger"`
	KafkaTopics      KafkaTopics          `mapstructure:"kafkaTopics"`
}

type Grpc struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

type MongoCollections struct {
	Users string `mapstructure:"users"`
}

type ServiceSettings struct {
	RedisProductFixKey string `mapstructure:"redisProductFixKey"`
}

type KafkaTopics struct {
	UserCreated productKafka.TopicConfig `mapstructure:"userCreated"`
	UserUpdated productKafka.TopicConfig `mapsstructure:"userUpdated"`
	UserDeleted productKafka.TopicConfig `mapstructure:"userDeleted"`
}

var configPath string

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/user_service/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	grpcPort := os.Getenv(constants.GrpcPort)
	if grpcPort != "" {
		cfg.GRPC.Port = grpcPort
	}
	postgresHost := os.Getenv(constants.PostgresqlHost)
	if postgresHost != "" {
		cfg.Postgres.Host = postgresHost
	}
	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		cfg.Postgres.Port = postgresPort
	}
	mongoURI := os.Getenv(constants.MongoDbURI)
	if mongoURI != "" {
		//cfg.Mongo.URI = "mongodb://host.docker.internal:27017"
		cfg.Mongo.Uri = mongoURI
	}
	kafkaBrokers := os.Getenv(constants.KafkaBrokers)
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}
	redisAddr := os.Getenv(constants.RedisAddr)
	if redisAddr != "" {
		cfg.Redis.Addr = redisAddr
	}
	jaegerAddr := os.Getenv(constants.JaegerHostPort)
	if jaegerAddr != "" {
		cfg.Jaeger.HostPort = jaegerAddr
	}

	return cfg, nil
}
