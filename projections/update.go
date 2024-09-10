package projections

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections/order"
	"github.com/nmarsollier/ordersgo/projections/status"
	"github.com/nmarsollier/ordersgo/tools/log"
)

func Update(orderId string, ctx ...interface{}) error {
	ev, err := events.FindByOrderId(orderId, ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return err
	}

	order, err := order.Update(orderId, ev, ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
	}

	status.Update(orderId, ev, order, ctx...)
	return nil
}
