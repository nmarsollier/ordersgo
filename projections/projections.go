package projections

import (
	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections/order"
	"github.com/nmarsollier/ordersgo/projections/status"
)

func Update(orderId string) error {
	ev, err := events.FindByOrderId(orderId)
	if err != nil {
		glog.Error(err)
		return err
	}

	order, err := order.UpdateProjection(orderId, ev)
	if err != nil {
		glog.Error(err)
	}

	status.UpdateProjection(orderId, ev, order)
	return nil
}
