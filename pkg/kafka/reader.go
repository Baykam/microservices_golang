package productKafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(kafkaUrl []string, topic, groupId string, errLogger kafka.Logger) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaUrl,
		GroupID:                groupId,
		Topic:                  topic,
		MinBytes:               MinBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		ErrorLogger:            errLogger,
		MaxAttempts:            maxAttempts,
		MaxWait:                time.Second,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	})
}
