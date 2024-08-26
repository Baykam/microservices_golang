package server

import (
	"context"
	"net/http"
	productKafka "project-microservices/pkg/kafka"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *server) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := productKafka.NewKafkaConn(ctx, s.cfg.Kafka)
	if err != nil {
		return errors.Wrap(err, "kafka.NewConnection")
	}
	s.kafka = kafkaConn

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return errors.Wrap(err, "kafkaConn.Brokers")
	}

	s.log.Infof("kafka connected to brokeres: %sv", brokers)

	return nil
}

func (s *server) getConsumerGroupTopics() []string {
	return []string{
		s.cfg.KafkaTopics.UserCreated.TopicName,
		s.cfg.KafkaTopics.UserDeleted.TopicName,
		s.cfg.KafkaTopics.UserUpdated.TopicName,
	}
}

func (s *server) runMetrics(cancel context.CancelFunc) {
	engine := gin.Default()
	engine.GET(s.cfg.Probes.PrometheusPath, gin.WrapH(promhttp.Handler()))
	go func() {
		s.log.Infof("Metrics server is listening : %s", s.cfg.Probes.PrometheusPort)
		ss := &http.Server{
			Addr:    s.cfg.Probes.PrometheusPort,
			Handler: engine,
		}
		if err := ss.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Errorf("metrics.Start:%v", err)
			cancel()
		}
	}()
}
