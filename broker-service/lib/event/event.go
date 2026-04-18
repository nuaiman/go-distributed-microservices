package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func getExchangeName() string {
	return "logs_topic"
}

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		getExchangeName(),
		"topic", // name
		true,    // durable
		false,   // auto-delete
		false,   // internal
		false,   // no-wait
		nil,     // args
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		true,  // durable
		false, // auto-delete
		true,  // exclusive
		false, // no-wait
		nil,   // args
	)
}
