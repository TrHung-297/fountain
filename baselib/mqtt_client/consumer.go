/* !!
 * File: consumer.go
 * File Created: Monday, 23rd August 2021 4:26:45 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Monday, 23rd August 2021 4:26:45 pm
 
 */

package mqtt_client

import (
	"encoding/json"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ConsumerProcessInstance interface {
	MQMessageError(err error)
	MQMessageCallback(messageObj MessageQueue)
}

func (manager *MQTTClient) InstallConsume(consumeInstance ConsumerProcessInstance, topic string) error {
	token := manager.Client.Subscribe(topic, KMqttDefaultQoS, func(client mqtt.Client, message mqtt.Message) {
		messageObj := MessageQueue{}
		if err := json.Unmarshal(message.Payload(), &messageObj); err != nil {
			consumeInstance.MQMessageError(err)
		} else {
			messageObj.Topic = message.Topic()
			consumeInstance.MQMessageCallback(messageObj)
		}
	})

	token.WaitTimeout(30 * time.Second)

	return token.Error()
}
