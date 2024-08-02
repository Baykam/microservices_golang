package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"project-microservices/pkg/constants"
	"project-microservices/pkg/logger"
	"project-microservices/pkg/probes"
	"project-microservices/pkg/tracing"

	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "API Gateway microservice config path")
}

type Config struct {
	ServiceName string          `mapstructure:"serviceName"`
	Logger      *logger.Config  `mapstructure:"logger"`
	Http        Http            `mapstructure:"http"`
	Grpc        Grpc            `mapstructure:"grpc"`
	Probes      probes.Config   `mapstructure:"probes"`
	Jaeger      *tracing.Config `mapstructure:"jaeger"`
	JWT         Jwt             `mapstructure:"jwt"`
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
			configPath = fmt.Sprintf("%s/api_gateway_service/config/config.yaml", getWd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Unwrap(err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Unwrap(err)
	}
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

	return cfg, nil
}
