package productKafka

type Config struct {
	Brokers    []string `mapstructure:"brokers"`
	GroupId    string   `mapstructure:"groupID"`
	InitTopics bool     `mapstructure:"initTopics"`
}

type TopicConfig struct {
	TopicName         string `mapstructure:"topicName"`
	Partitions        int    `mapstructure:"partitions"`
	ReplicationFactor int    `mapstructure:"replicationFactor"`
}
