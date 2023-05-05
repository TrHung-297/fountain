

package amqp_client

import (
	"github.com/streadway/amqp"
	"github.com/TrHung-297/fountain/baselib/base"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

func (client *AmqpClient) GetChannel() (*amqp.Channel, error) {
	return client.Channel()
}

func (client *AmqpClient) CreateQueue(name string) (amqp.Queue, error) {
	ch, err := client.Channel()
	if err != nil {
		g_log.V(1).Errorf("AmqpClient::CreateQueue - Can not get channel - Error: %+v", err)

		return amqp.Queue{}, err
	}

	defer ch.Close()
	return ch.QueueDeclare(
		name,  // name of the queue
		true,  // should the message be persistent? also queue will survive if the cluster gets reset
		false, // auto delete if there's no consumers (like queues that have anonymous names, often used with fanout exchange)
		false, // exclusive means I should get an error if any other consumer subscribes to this queue
		false, // no-wait means I don't want RabbitMQ to wait if there's a queue successfully setup
		amqp.Table{"x-queue-type": config.QueueType}, // arguments for more advanced configuration
	)
}

func (client *AmqpClient) QueueBind(queueName, exchange string) error {
	ch, err := client.Channel()
	if err != nil {
		g_log.V(1).Errorf("AmqpClient::CreateQueue - Can not get channel - Error: %+v", err)

		return err
	}

	defer ch.Close()
	return ch.QueueBind(queueName, queueName, exchange, false, nil)
}

func (client *AmqpClient) PushMessage(exchange string, routeKey string, message *MessageQueue) error {
	channel, err := client.Channel()
	if err != nil {
		g_log.V(1).Errorf("AmqpClient::PushMessage - Can not get channel - Error: %+v", err)

		return err
	}
	defer channel.Close()

	// Set RabbitMQ QoS
	if err = channel.Qos(1, 0, false); err != nil {
		g_log.V(1).Errorf("AmqpClient::PushMessage - Error Set Qos Channel %+v", err)

		return err
	}

	err = channel.Publish(
		exchange, // exchange
		routeKey, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  KDefaultContentType,
			Body:         base.JSONDebugData(message),
		})

	if err != nil {
		g_log.V(1).Errorf("AmqpClient::PushMessage - Failed to publish a message %+v", err)

		return err
	}

	return nil
}
