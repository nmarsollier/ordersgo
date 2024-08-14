package order_projection

import (
	"github.com/nmarsollier/ordersgo/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateOrderProjection(orderId string) error {
	ev, err := events.FindByOrderId(orderId)
	if err != nil {
		return err
	}

	order, _ := FindByOrderId(orderId)
	if order == nil {
		order = &Order{
			ID:      primitive.NewObjectID(),
			OrderId: orderId,
		}
	}

	for _, e := range ev {
		order = order.update(e)
	}

	if _, err := insert(order); err != nil {
		return err
	}

	return nil
}

func (order *Order) update(event *events.Event) *Order {
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

func (o *Order) updadatePlace(e *events.Event) *Order {
	o.OrderId = e.OrderId
	o.UserId = e.PlaceEvent.UserId
	o.CartId = e.PlaceEvent.CartId
	o.Status = Placed
	o.Created = e.Created
	o.Updated = e.Updated

	articles := make([]*Article, len(e.PlaceEvent.Articles))
	for i, article := range e.PlaceEvent.Articles {
		articles[i] = &Article{
			ArticleId: article.ArticleId,
			Quantity:  article.Quantity,
		}
	}

	o.Articles = articles
	return o
}

func (o *Order) updadateValidation(e *events.Event) *Order {
	validation := e.Validation

	for _, a := range o.Articles {
		if a.ArticleId == validation.ArticleId {
			a.IsValid = validation.IsValid
			a.UnitaryPrice = validation.Price
			a.IsValidated = true
		}
	}

	o.Status = Validated
	for _, a := range o.Articles {
		if !a.IsValid {
			o.Status = Invalid
		}
	}

	o.Updated = e.Updated

	return o
}

func (o *Order) updadatePayment(e *events.Event) *Order {
	o.Payments = append(o.Payments, &PaymentEvent{
		Method: e.Payment.Method,
		Amount: e.Payment.Amount,
	})

	if o.TotalPayment() >= o.TotalPrice() {
		o.Status = Payment_Defined
	}

	o.Updated = e.Updated

	return o
}
