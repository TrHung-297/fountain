/* !!
 * File: consume.go
 * File Created: Thursday, 24th June 2021 3:22:40 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 24th June 2021 3:22:40 pm
 
 */

package amqp_client

import (
	"encoding/json"
	"time"

	"github.com/streadway/amqp"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
)

type ConsumerProcessInstance interface {
	AMMessageError(err error)
	AMMessageCallback(messageObj MessageQueue)
}

func (client *AmqpClient) InstallConsume(consumeExchange, consumeQueue string, consumeInstance ConsumerProcessInstance) error {
	channel, err := client.Channel()
	if err != nil {
		g_log.V(1).Errorf("AmqpClient::PushMessage - Can not get channel - Error: %+v", err)

		return err
	}

	if err := channel.ExchangeDeclare(consumeExchange, amqp.ExchangeDirect, true, false, false, false, nil); err != nil {
		g_log.V(1).Errorf("AmqpClient::InstallConsume - QueueDeclare Error: %v", err)
		return err
	}

	g_log.V(1).Infof("AmqpClient::InstallConsume - Create queue: %s for exchange: %s", consumeQueue, consumeExchange)
	queue, err := channel.QueueDeclare(
		consumeQueue, // name of the queue
		true,         // should the message be persistent? also queue will survive if the cluster gets reset
		false,        // auto delete if there's no consumers (like queues that have anonymous names, often used with fanout exchange)
		false,        // exclusive means I should get an error if any other consumer subscribes to this queue
		false,        // no-wait means I don't want RabbitMQ to wait if there's a queue successfully setup
		amqp.Table{"x-queue-type": config.QueueType}, // arguments for more advanced configuration
	)

	if err != nil {
		g_log.V(1).Errorf("AmqpClient::InstallConsume - QueueDeclare Error: %v", err)
		return err
	}

	err = channel.QueueBind(consumeQueue, consumeQueue, consumeExchange, false, nil)
	if err != nil {
		g_log.V(1).Errorf("AmqpClient::InstallConsume - QueueBind Error: %v", err)
		return err
	}

	messages, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		g_log.V(1).Errorf("AmqpClient::InstallConsume - QueueBind Error: %v", err)
		return err
	}

	go func() {
		for message := range messages {
			eventTask := MessageQueue{}
			if err := json.Unmarshal(message.Body, &eventTask); err != nil {
				consumeInstance.AMMessageError(err)
				message.Ack(false)
				continue
			}

			consumeInstance.AMMessageCallback(eventTask)
			message.Ack(false)
			time.Sleep(5 * time.Millisecond)
		}
	}()

	return nil
}
