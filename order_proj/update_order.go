package order_proj

import "github.com/nmarsollier/ordersgo/events"

func UpdateOrderProjection(orderId string) error {

	ev, err := events.FindByOrderId(orderId)

	if err != nil {
		return err
	}

	order := Order{}
	for _, e := range ev {
		order = order.Update(e)
	}

	insert(&order)

	return nil
}

func (order Order) Update(event *events.Event) Order {
	switch event.Type {
	case events.Place:
		order = order.UpdadatePlace(event)
	case events.Validation:
		order = order.UpdadateValidation(event)
	case events.Payment:
		order = order.UpdadatePayment(event)
	}
	return order
}

func (o Order) UpdadatePlace(e *events.Event) Order {
	o.OrderId = e.OrderId
	o.UserId = e.PlaceEvent.UserId
	o.CartId = e.PlaceEvent.CartId
	o.Status = Placed

	articles := make([]Article, len(e.PlaceEvent.Articles))
	for i, article := range e.PlaceEvent.Articles {
		articles[i] = Article{
			ArticleId: article.ArticleId,
			Quantity:  article.Quantity,
		}
	}

	o.Articles = articles
	return o
}

func (o Order) UpdadateValidation(e *events.Event) Order {
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

	return o
}

func (o Order) UpdadatePayment(e *events.Event) Order {
	o.Payments = append(o.Payments, PaymentEvent{
		Method: e.Payment.Method,
		Amount: e.Payment.Amount,
	})

	if o.TotalPayment() >= o.TotalPrice() {
		o.Status = Payment_Defined
	}
	return o
}
