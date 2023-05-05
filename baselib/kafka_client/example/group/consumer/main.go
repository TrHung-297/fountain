

package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/base"
	"github.com/TrHung-297/fountain/baselib/kafka_client"
)

func init() {
	viper.Set("Kafka.Broker", "kafka-chatting.gtvplusdev.info.private:9092")
	viper.Set("Kafka.ProducerTopics", "gtv_test")
	viper.Set("Kafka.ConsumerGroupName", "gtv_consumer_test")
	viper.Set("Kafka.ConsumerTopicNames", "gtv_test")
}

func main() {
	kafka_client.InstallKafkaClient()

	kafClient := kafka_client.GetKafkaClientInstance()

	consumerCallback := &ConsumerInstace{}
	kafClient.InstallConsumerGroup(consumerCallback, "gtv_test")

	fmt.Printf("Consumer done!\n")
}

// ConsumerInstace func;
type ConsumerInstace struct{}

// ErrorCallback func;
func (c *ConsumerInstace) ErrorCallback(err error) {
	fmt.Printf("errCallback::Error - %+v\n", err)
}

// MessageCallback func;
func (c *ConsumerInstace) MessageCallback(msgObj kafka_client.MessageKafka) {
	fmt.Printf("procCallback::msgObj - %s\n", base.JSONDebugDataString(msgObj))
}
