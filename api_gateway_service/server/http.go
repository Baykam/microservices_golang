package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *server) runHttpServer() {
	s.engine.SetTrustedProxies(nil)

	// s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	s.engine.Use(gin.Recovery())

	if err := s.engine.Run(s.cfg.Http.Port); err != nil {
		log.Fatalf("Running http server error: %v", err)
	}
}
