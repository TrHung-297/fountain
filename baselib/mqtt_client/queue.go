

package mqtt_client

import (
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
)

const (
	KMqttDefaultQoS = 0
	// MqttDefaultQuiesce defines default Quiesce
	KMqttDefaultQuiesce = 250
)

/**
 * Publish Message
 */
func (manager *MQTTClient) Publish(topic string, message MessageQueue) {
	token := manager.Client.Publish(topic, KMqttDefaultQoS, false, base.JSONDebugDataString(message))
	token.Wait()
	err := token.Error()
	if err != nil {
		g_log.V(1).WithError(err).Errorf("MQTTClient::Publish - Publish to mqtt error: %s, topic: %s, msg: %s", err, topic, base.JSONDebugDataString(message))
	}
}
