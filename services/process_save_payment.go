package services

import (
	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/order_projection"
)

func ProcessSavePayment(data *events.PaymentEvent) (*events.Event, error) {
	event, err := events.SavePayment(data)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	go order_projection.UpdateOrderProjection(event.OrderId)

	return event, err
}
