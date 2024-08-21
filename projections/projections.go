package projections

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/log"
	"github.com/nmarsollier/ordersgo/projections/order"
	"github.com/nmarsollier/ordersgo/projections/status"
)

func Update(orderId string, ctx ...interface{}) error {
	ev, err := events.FindByOrderId(orderId, ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
		return err
	}

	order, err := order.UpdateProjection(orderId, ev, ctx...)
	if err != nil {
		log.Get(ctx...).Error(err)
	}

	status.UpdateProjection(orderId, ev, order, ctx...)
	return nil
}
