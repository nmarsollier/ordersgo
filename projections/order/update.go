package order

import (
	"github.com/nmarsollier/ordersgo/events"
)

func Update(orderId string, ev []*events.Event, ctx ...interface{}) (*Order, error) {
	order, _ := FindByOrderId(orderId, ctx...)
	if order == nil {
		order = &Order{
			OrderId: orderId,
		}
	}

	for _, e := range ev {
		order = order.update(e)
	}

	if _, err := insert(order, ctx...); err != nil {
		return nil, err
	}

	return order, nil
}

func (order *Order) update(event *events.Event) *Order {
	switch event.Type {
	case events.Place:
		order = order.updatePlace(event)
	case events.Validation:
		order = order.updateValidation(event)
	case events.Payment:
		order = order.updatePayment(event)
	}
	return order
}

func (o *Order) updatePlace(e *events.Event) *Order {
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

func (o *Order) updateValidation(e *events.Event) *Order {
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

func (o *Order) updatePayment(e *events.Event) *Order {
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
