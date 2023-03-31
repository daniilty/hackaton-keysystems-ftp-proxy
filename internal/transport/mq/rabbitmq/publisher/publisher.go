package publisher

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	SendContract([]byte) error
}

type publisher struct {
	channel  *amqp.Channel
	exchange string
}

func NewPublisher(channel *amqp.Channel, exchange string) Publisher {
	return &publisher{
		channel:  channel,
		exchange: exchange,
	}
}

func (p *publisher) SendContract(data []byte) error {
	const contentType = "application/json"

	return p.channel.Publish(p.exchange, "*", false, false, amqp.Publishing{
		Headers: amqp.Table{
			"content-type": contentType,
		},
		Body: data,
	})
}
