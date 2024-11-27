package projections

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections/order"
	"github.com/nmarsollier/ordersgo/projections/status"
	"github.com/nmarsollier/ordersgo/tools/log"
)

func Update(orderId string, deps ...interface{}) error {
	ev, err := events.FindByOrderId(orderId, deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	order, err := order.Update(orderId, ev, deps...)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	status.Update(orderId, ev, order, deps...)
	return nil
}
