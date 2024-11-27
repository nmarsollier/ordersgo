package emit

import (
	"errors"

	"github.com/nmarsollier/ordersgo/tools/env"
	"github.com/nmarsollier/ordersgo/tools/log"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.New("channel not initialized")

func getChannel(deps ...interface{}) (*amqp.Channel, error) {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}
	if ch == nil {
		return nil, ErrChannelNotInitialized
	}
	return ch, nil
}
