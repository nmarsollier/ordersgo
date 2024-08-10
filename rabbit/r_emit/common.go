package r_emit

import (
	"github.com/nmarsollier/ordersgo/tools/env"
	"github.com/nmarsollier/ordersgo/tools/errors"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.NewCustom(400, "Channel not initialized")

func getChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	if ch == nil {
		return nil, ErrChannelNotInitialized
	}
	return ch, nil
}
