package producer

import "github.com/IBM/sarama"

func initKafkaConfig() *sarama.Config {
	prodCfg := sarama.NewConfig()
	prodCfg.Producer.Partitioner = sarama.NewRandomPartitioner
	prodCfg.Producer.RequiredAcks = sarama.WaitForAll
	prodCfg.Producer.Return.Successes = true
	return prodCfg
}
