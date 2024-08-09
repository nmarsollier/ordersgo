package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/order_proj"
)

func PocessPlaceOrder(data *events.PlacedOrderData) (*events.Event, error) {
	event, err := events.SavePlaceOrder(data)
	if err != nil {
		return nil, err
	}

	go order_proj.UpdateOrderProjection(event.OrderId)

	return event, err
}
