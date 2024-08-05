package main

import (
	"flag"
	"project-microservices/api_gateway_service/config"
	"project-microservices/api_gateway_service/server"
	_ "project-microservices/docs"
	"project-microservices/pkg/logger"

	"github.com/pkg/errors"
)

// @title API Gateway Service
// @version 1.0
// @description API Gateway Service için Swagger dokümantasyonu
func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		errors.Wrap(err, "InitConfig.ApiGatewayService")
	}

	appLogger := logger.NewAppLogger(cfg.Logger)

	sv := server.NewServer(appLogger, cfg)
	if err := sv.Run(); err != nil {
		errors.Wrap(err, "Run.ApiGatewayService")
	}
}
