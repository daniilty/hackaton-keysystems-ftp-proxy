package publisher

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	SendContract([]byte) error
}

type publisher struct {
	channel *amqp.Channel
}

func NewPublisher(channel *amqp.Channel) Publisher {
	return &publisher{
		channel: channel,
	}
}

func (p *publisher) SendContract(data []byte) error {
	const (
		exchange    = "contracts.changes"
		contentType = "application/json"
	)

	return p.channel.Publish(exchange, "*", false, false, amqp.Publishing{
		Headers: amqp.Table{
			"content-type": contentType,
		},
		Body: data,
	})
}
