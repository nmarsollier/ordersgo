package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections"
	"github.com/nmarsollier/ordersgo/tools/log"
)

func ProcessSavePayment(data *events.PaymentEvent, deps ...interface{}) (*events.Event, error) {
	event, err := events.SavePayment(data, deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	go projections.Update(event.OrderId, deps...)

	return event, err
}
