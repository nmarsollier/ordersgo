package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/order_projection"
)

func ProcessArticleData(data *events.ValidationEvent) (*events.Event, error) {
	event, err := events.SaveArticleExist(data)
	if err != nil {
		return nil, err
	}

	go order_projection.UpdateOrderProjection(event.OrderId)

	return event, err
}
