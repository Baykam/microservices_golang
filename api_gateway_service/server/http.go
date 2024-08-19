package server

import (
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "project-microservices/docs"
)

func (s *server) runHttpServer() {
	s.engine.SetTrustedProxies(nil)

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := s.engine.Run(s.cfg.Http.Port); err != nil {
		log.Fatalf("Running http server error: %v", err)
	}
}
