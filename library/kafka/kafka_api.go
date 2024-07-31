package kafka

import "github.com/Shopify/sarama"

type Kafka_Client struct {
	client *sarama.Client
}

func NewKafka_Client(config *ClientConfig) (*Kafka_Client, error) {
	conf := sarama.NewConfig()
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Metadata.Full = true
	conf.Version = sarama.V0_10_0_0
	conf.Net.SASL.Enable = true
	conf.Net.SASL.User = config.UserName
	conf.Net.SASL.Password = config.Password
	conf.Net.SASL.Handshake = true
	if config.mode == "consume" {
	}
	return nil, nil
}
