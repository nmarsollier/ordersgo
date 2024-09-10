package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections"
	"github.com/nmarsollier/ordersgo/tools/log"
)

func ProcessSavePayment(data *events.PaymentEvent, ctx ...interface{}) (*events.Event, error) {
	event, err := events.SavePayment(data, ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	go projections.Update(event.OrderId, ctx...)

	return event, err
}
