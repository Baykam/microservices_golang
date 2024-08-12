package config

import (
	"errors"
	"fmt"
	"os"
	"project-microservices/pkg/constants"
	productKafka "project-microservices/pkg/kafka"
	"project-microservices/pkg/logger"
	"project-microservices/pkg/probes"
	"project-microservices/pkg/tracing"
	"strings"

	"github.com/spf13/viper"
)

var configPath string

type Config struct {
	ServiceName string               `mapstructure:"serviceName"`
	Logger      *logger.Config       `mapstructure:"logger"`
	Http        Http                 `mapstructure:"http"`
	Grpc        Grpc                 `mapstructure:"grpc"`
	Probes      probes.Config        `mapstructure:"probes"`
	Jaeger      *tracing.Config      `mapstructure:"jaeger"`
	JWT         Jwt                  `mapstructure:"jwt"`
	KafkaTopics KafkaTopics          `mapstructure:"kafkaTopics"`
	Kafka       *productKafka.Config `mapstructure:"kafka"`
}

type Http struct {
	Port                string   `mapstructure:"port"`
	Development         bool     `mapstructure:"development"`
	BasePath            string   `mapstructure:"basePath"`
	ProductsPath        string   `mapstructure:"productsPath"`
	DebugHeaders        bool     `mapstructure:"debugHeaders"`
	HttpClientDebug     bool     `mapstructure:"httpClientDebug"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
}

type Grpc struct {
	UserServicePort string `mapstructure:"user_service"`
}

type Jwt struct {
	SecretKey string `mapstructure:"secretKey"`
}

type KafkaTopics struct {
	UserCreated productKafka.TopicConfig `mapstructure:"userCreated"`
	UserUpdated productKafka.TopicConfig `mapstructure:"userUpdated"`
	UserDeleted productKafka.TopicConfig `mapstructure:"userDeleted"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getWd, err := os.Getwd()
			if err != nil {
				return nil, errors.Unwrap(err)
			}
			trimmedWd := strings.TrimSuffix(getWd, "/cmd")
			configPath = fmt.Sprintf("%s/config/config.yaml", trimmedWd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Unwrap(err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Unwrap(err)
	}

	if err := setEnvOverrides(cfg); err != nil {
		return nil, fmt.Errorf("error setting environment overrides: %w", err)
	}

	return cfg, nil
}

func setEnvOverrides(cfg *Config) error {
	secretKey := os.Getenv(constants.SecretKey)
	if secretKey != "" {
		cfg.JWT.SecretKey = secretKey
	}

	httpPort := os.Getenv(constants.HttpPort)
	if httpPort != "" {
		cfg.Http.Port = httpPort
	}

	jaegerAddr := os.Getenv(constants.JaegerHostPort)
	if jaegerAddr != "" {
		cfg.Jaeger.HostPort = jaegerAddr
	}

	readerServicePort := os.Getenv(constants.UserService)
	if readerServicePort != "" {
		cfg.Grpc.UserServicePort = readerServicePort
	}

	return nil
}
