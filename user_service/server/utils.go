package server

// func (s *server) connectKafkaBrokers(ctx context.Context) error {
// 	kafkaConn, err := productKafka.NewKafkaConn(ctx, s.cfg.Kafka)
// 	if err != nil {
// 		return errors.Wrap(err, "kafka.NewConnection")
// 	}
// 	s.kafka = kafkaConn

// 	brokers, err := kafkaConn.Brokers()
// 	if err != nil {
// 		return errors.Wrap(err, "kafkaConn.Brokers")
// 	}

// 	s.log.Infof("kafka connected to brokeres: %sv", brokers)

// 	return nil
// }

func (s *server) getConsumerGroupTopics() []string {
	return []string{
		s.cfg.KafkaTopics.UserCreated.TopicName,
		s.cfg.KafkaTopics.UserDeleted.TopicName,
		s.cfg.KafkaTopics.UserUpdated.TopicName,
	}
}
