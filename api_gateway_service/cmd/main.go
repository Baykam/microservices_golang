package main

import (
	"flag"
	"project-microservices/api_gateway_service/config"
	"project-microservices/api_gateway_service/server"

	"project-microservices/pkg/logger"

	"github.com/pkg/errors"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		errors.Wrap(err, "InitConfig.ApiGatewayService")
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()

	sv := server.NewServer(appLogger, cfg)
	if err := sv.Run(); err != nil {
		errors.Wrap(err, "Run.ApiGatewayService")
	}
}
