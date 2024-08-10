package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/order_proj"
	"github.com/nmarsollier/ordersgo/rabbit/r_emit"
)

func PocessPlaceOrder(data *events.PlacedOrderData) (*events.Event, error) {
	event, err := events.SavePlaceOrder(data)
	if err != nil {
		return nil, err
	}

	go order_proj.UpdateOrderProjection(event.OrderId)

	r_emit.EmitOrderPlaced(event)

	for _, article := range event.PlaceEvent.Articles {
		go r_emit.EmitArticleValidation(r_emit.ArticleValidationData{
			ReferenceId: event.OrderId,
			ArticleId:   article.ArticleId,
		})
	}

	return event, err
}
