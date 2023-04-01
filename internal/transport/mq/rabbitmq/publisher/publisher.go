package publisher

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	SendContract(context.Context, []byte) error
	SendContractProcedure(context.Context, []byte) error
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

func (p *publisher) SendContract(ctx context.Context, data []byte) error {
	const contentType = "application/json"

	return p.channel.PublishWithContext(ctx, p.exchange, "contract", false, false, amqp.Publishing{
		Headers: amqp.Table{
			"content-type": contentType,
		},
		Body: data,
	})
}

func (p *publisher) SendContractProcedure(ctx context.Context, data []byte) error {
	const contentType = "application/json"

	return p.channel.PublishWithContext(ctx, p.exchange, "contract-procedure", false, false, amqp.Publishing{
		Headers: amqp.Table{
			"content-type": contentType,
		},
		Body: data,
	})
}
