/* !!
 * File: client.go
 * File Created: Wednesday, 9th June 2021 3:23:13 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 9th June 2021 3:23:13 pm
 
 */

package amqp_client

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

type AmqpClient struct {
	*amqp.Connection
}

var amInstance *AmqpClient

// default value env key is "AMQP";
// if configKeys was set, key env will be first value (not empty) of this;
func InstanceAMQPClientManager(configKeys ...string) *AmqpClient {
	if amInstance != nil {
		return amInstance
	}

	if config == nil {
		getAmqpConfigFromEnv(configKeys...)
	}

	if config == nil {
		err := fmt.Errorf("need config for amqp client first")
		panic(err)
	}

	retry := 0
	for amInstance == nil && retry < 3 {
		retry++

		url := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port)

		conn, err := amqp.Dial(url)
		if err != nil {
			g_log.V(1).Errorf("Failed to connect to AMQP - Error: %+v", err)
			continue
		}

		// Reconnect to AMQP in case connection died
		time.Sleep(500 * time.Millisecond)

		amInstance = &AmqpClient{conn}

		go func() {
			for {
				rabbitCloseError := make(chan *amqp.Error)
				amInstance.NotifyClose(rabbitCloseError)

				if rabbitErr, ok := <-rabbitCloseError; !ok {
					g_log.V(1).Errorf("connection to AMQP was closed, do not need try to reconnect")
					return
				} else if rabbitErr != nil {
					g_log.V(1).Errorf("trying to connect to AMQP - Error: %+v", rabbitErr)
				}
			}
		}()
	}

	if amInstance == nil {
		err := fmt.Errorf("can not create instance for amqp")
		panic(err)
	}

	return amInstance
}

func GetAmqpClientInstance() *AmqpClient {
	if amInstance == nil {
		return InstanceAMQPClientManager()
	}

	return amInstance
}
