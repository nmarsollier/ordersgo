package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/order_proj"
)

func ProcessSavePayment(data *events.PaymentEvent) (*events.Event, error) {
	event, err := events.SavePayment(data)

	go order_proj.UpdateOrderProjection(event.OrderId)

	return event, err
}
