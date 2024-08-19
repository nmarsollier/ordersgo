package status

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections/order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProjection(orderId string, ev []*events.Event, odr *order.Order) error {
	status, _ := FindByOrderId(orderId)
	if status == nil {
		status = &OrderStatus{
			ID:      primitive.NewObjectID(),
			OrderId: orderId,
		}
	}

	switch odr.Status {
	case order.Validated:
		{
			status.Validated = true
		}
	case order.Payment_Defined:
		{
			status.PaymentCompleted = true
		}
	}

	for _, e := range ev {
		status = status.update(e)
	}

	if _, err := insert(status); err != nil {
		return err
	}

	return nil
}

func (order *OrderStatus) update(event *events.Event) *OrderStatus {
	switch event.Type {
	case events.Place:
		order = order.updadatePlace(event)
	case events.Validation:
		order = order.updadateValidation(event)
	case events.Payment:
		order = order.updadatePayment(event)
	}
	return order
}

func (o *OrderStatus) updadatePlace(e *events.Event) *OrderStatus {
	o.OrderId = e.OrderId
	o.UserId = e.PlaceEvent.UserId
	o.Placed = true

	o.Created = e.Created
	o.Updated = e.Updated
	return o
}

func (o *OrderStatus) updadateValidation(e *events.Event) *OrderStatus {
	if e.Validation.IsValid {
		o.PartialValidated = true
	}
	o.Updated = e.Updated
	return o
}

func (o *OrderStatus) updadatePayment(e *events.Event) *OrderStatus {
	if e.Payment.Amount > 0 {
		o.PartialPayment = true
	}
	o.Updated = e.Updated
	return o
}
