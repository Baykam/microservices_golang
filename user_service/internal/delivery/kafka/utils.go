package userKafkaConn

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func (s *userMessageProcessor) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	// s.metrics.SuccessKafkaMessages.Inc()
	s.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		s.log.WarnMsg("commitMessage", err)
	}
}
