package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/order_proj"
)

func ProcessArticleData(data *events.ValidationEvent) (*events.Event, error) {
	event, err := events.SaveArticleExist(data)

	go order_proj.UpdateOrderProjection(event.OrderId)

	return event, err
}
