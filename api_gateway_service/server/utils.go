package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	readTimeout  = 15 * time.Second
	writeTimeout = 15 * time.Second
)

func (s *server) runHealthCheck(ctx context.Context) {
	health := healthcheck.NewHandler()

	health.AddReadinessCheck(s.cfg.ServiceName, healthcheck.AsyncWithContext(ctx, func() error {
		if s.cfg != nil {
			return nil
		}
		return errors.New("config not load")
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))

	go func() {
		s.log.Infof("Api gateway Kubernetes probes listening on port: %s", s.cfg.Probes.Port)
		if err := http.ListenAndServe(s.cfg.Probes.Port, health); err != nil {
			s.log.WarnMsg("ListenAndServe", err)
		}
	}()
}

func (s *server) runMetrics(cancel context.CancelFunc) {
	s.engine.Use(gin.Recovery())
	s.engine.GET(s.cfg.Probes.PrometheusPath, gin.WrapH(promhttp.Handler()))
	go func() {
		s.log.Infof("Metrics server is listening : %s", s.cfg.Probes.PrometheusPort)
		ss := &http.Server{
			Addr:         s.cfg.Probes.PrometheusPort,
			Handler:      s.engine,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		}
		if err := ss.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Errorf("metrics.Start:%v", err)
			cancel()
		}
	}()
}
