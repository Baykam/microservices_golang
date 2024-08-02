package main

import (
	"flag"
	"project-microservices/api_gateway_service/config"
	"project-microservices/api_gateway_service/server"
	"project-microservices/pkg/logger"

	_ "project-microservices/api_gateway_service/docs/user"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	_ "github.com/swaggo/swag"
)

// @title Tag Service Api
// @Description hey there
func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		errors.Wrap(err, "InitConfig.ApiGatewayService")
	}

	appLogger := logger.NewAppLogger(cfg.Logger)

	engine := gin.Default()

	sv := server.NewServer(appLogger, cfg, engine)
	if err := sv.Run(); err != nil {
		errors.Wrap(err, "Run.ApiGatewayService")
	}
}
