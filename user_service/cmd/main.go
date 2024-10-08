package main

import (
	"flag"
	"project-microservices/pkg/logger"
	"project-microservices/user_service/config"
	"project-microservices/user_service/server"

	"github.com/pkg/errors"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		errors.Wrap(err, "InitConfig.UserService")
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()

	sv := server.NewServer(appLogger, cfg)
	if err := sv.Run(); err != nil {
		errors.Wrap(err, "Run.UserService")
	}
}
