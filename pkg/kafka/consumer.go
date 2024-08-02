package productKafka

import (
	"context"
	"project-microservices/pkg/logger"
	"sync"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

type ConsumerGroup interface {
	GetNewKafkaReader(kafkaURL []string, groupTopics []string, groupID string) *kafka.Reader
	GetNewKafkaWriter() *kafka.Writer
	ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker)
}

// Worker kafka consumer worker fetch and process messages from reader
type Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

type consumerGroup struct {
	Brokers []string
	GroupId string
	log     logger.Logger
}

func NewConsumerGroup(brokers []string,
	groupId string,
	log logger.Logger) *consumerGroup {
	return &consumerGroup{log: log, Brokers: brokers, GroupId: groupId}
}

// GetNewKafkaReader create new kafka reader
func (c *consumerGroup) GetNewKafkaReader(kafkaURL []string, groupTopics []string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		GroupTopics:            groupTopics,
		MinBytes:               MinBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            maxAttempts,
		MaxWait:                maxWait,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	})
}

// GetNewKafkaWriter create new kafka producer
func (c *consumerGroup) GetNewKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(c.Brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: writerRequiredAcks,
		MaxAttempts:  writerMaxAttempts,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
	}
}

// ConsumeTopic start consumer group with given worker and pool size
func (c *consumerGroup) ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker) {
	r := c.GetNewKafkaReader(c.Brokers, groupTopics, c.GroupId)

	defer func() {
		if err := r.Close(); err != nil {
			c.log.Warnf("consumerGroup.r.Close: %v", err)
		}
	}()
	c.log.Infof("Starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupId, groupTopics, poolSize)

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}
