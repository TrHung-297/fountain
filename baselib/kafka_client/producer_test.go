

package kafka_client

import (
	"math/rand"
	"testing"

	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/base"
)

func init() {
	viper.Set("Kafka.Broker", "kafka-chatting.gtvplusdev.info.private:9092")
	viper.Set("Kafka.ProducerTopics", "gtv_test")
	viper.Set("Kafka.ConsumerGroupName", "gtv_consumer_test")
	viper.Set("Kafka.ConsumerTopicNames", "gtv_test")
}

// Test that TestInstallKafkaClient works.
func TestInstallKafkaClient(t *testing.T) {
	instance := InstallKafkaClient()
	if instance == nil {
		t.Errorf("TestInstallKafkaClient - Error can not create kafka instance")
	}
}

type DataTmp struct {
	Data int64 `json:"data,omitempty"`
}

// Test that shortHostname works as advertised.
func TestProducerPushMessage(t *testing.T) {
	instance := InstallKafkaClient()
	if instance == nil {
		t.Errorf("TestProducerPushMessage - Error can not create kafka instance")
	}

	randInt := rand.Int63()
	tmp := DataTmp{
		Data: randInt,
	}
	msg := MessageKafka{
		Event:      "Test",
		ObjectJSON: base.JSONDebugDataString(tmp),
	}

	par, off, err := instance.ProducerPushMessage("gtv_test", msg)
	if err != nil {
		t.Errorf("TestProducerPushMessage - ProducerPushMessage Error %+v while result expect nil", err)
	}

	if off == 0 {
		t.Errorf("TestProducerPushMessage - ProducerPushMessage offset is 0 while result expect greater 0")
	}

	t.Logf("Partion: %d", par)
}

type NullService struct {
	cbErr func(err error)
	cbMsg func(msg MessageKafka)
}

func (*NullService) ErrorCallback(err error) {

}
func (*NullService) MessageCallback(messageObj MessageKafka) {

}

// Test that TestKafkaConsumer works
func TestKafkaConsumer(t *testing.T) {
	instance := InstallKafkaClient()
	if instance == nil {
		t.Errorf("TestKafkaConsumer - Error can not create kafka instance")
	}

	randInt := rand.Int63()
	tmp := DataTmp{
		Data: randInt,
	}

	dataJSON := base.JSONDebugDataString(tmp)
	msg := MessageKafka{
		Event:      "Test",
		ObjectJSON: dataJSON,
	}

	nullInstance := &NullService{
		cbErr: func(err error) {
			if err != nil {
				t.Errorf("TestKafkaConsumer - Error is %+v while expect nil", err)
			}
		},

		cbMsg: func(msg MessageKafka) {
			if msg.Event != "Test1" {
				t.Errorf("TestKafkaConsumer - Message event is %q while expect %q", msg.Event, "Test")
			}

			if msg.ObjectJSON == "123" {
				t.Errorf("TestKafkaConsumer - Message json data is %q while expect %q", msg.ObjectJSON, dataJSON)
			}
		},
	}

	go instance.InstallConsumerGroup(nullInstance, "gtv_test")

	par, off, err := instance.ProducerPushMessage("gtv_test", msg)
	if err != nil {
		t.Errorf("TestProducerPushMessage - ProducerPushMessage Error %+v while result expect nil", err)
	}

	if off == 0 {
		t.Errorf("TestProducerPushMessage - ProducerPushMessage offset is 0 while result expect greater 0")
	}

	_ = par
}

func BenchmarkHeader(b *testing.B) {
	instance := InstallKafkaClient()
	if instance == nil {
		b.Errorf("BenchmarkHeader - Error can not create kafka instance")
	}

	randInt := rand.Int63()
	tmp := DataTmp{
		Data: randInt,
	}

	dataJSON := base.JSONDebugDataString(tmp)
	msg := MessageKafka{
		Event:      "Test",
		ObjectJSON: dataJSON,
	}
	for i := 0; i < b.N; i++ {
		instance.ProducerPushMessage("gtv_test", msg)
	}
}
