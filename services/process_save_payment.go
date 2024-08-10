package services

import (
	"fmt"

	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/order_proj"
)

func ProcessSavePayment(data *events.PaymentEvent) (*events.Event, error) {
	event, err := events.SavePayment(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	go order_proj.UpdateOrderProjection(event.OrderId)

	return event, err
}