package rabbit

import (
	"github.com/nmarsollier/ordersgo/tools/env"
	"github.com/streadway/amqp"
)

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

type ConsumeMessage struct {
	Type     string `json:"type"`
	Version  int    `json:"version"`
	Queue    string `json:"queue"`
	Exchange string `json:"exchange"`
	Message  string `json:"message"`
}
