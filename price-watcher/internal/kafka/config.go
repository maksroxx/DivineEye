package kafka

import "github.com/IBM/sarama"

type ConfigProducer struct {
	Brokers []string
	Topics  []string
}

func NewProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Version = sarama.V2_8_0_0
	return config
}
