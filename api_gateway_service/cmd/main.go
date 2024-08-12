package main

import (
	"flag"
	"project-microservices/api_gateway_service/config"
	"project-microservices/api_gateway_service/server"

	"project-microservices/pkg/logger"

	_ "project-microservices/docs"

	"github.com/pkg/errors"
)

// @title Tag Service API
// @version 1.0
// @description A Tag service API in GO using gin framework

// @host 127.0.0.1:5001
// @BasePath  /api
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
