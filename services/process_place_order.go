package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/order_projection"
	"github.com/nmarsollier/ordersgo/rabbit/emit"
)

func PocessPlaceOrder(data *events.PlacedOrderData) (*events.Event, error) {
	event, err := events.SavePlaceOrder(data)
	if err != nil {
		return nil, err
	}

	go order_projection.UpdateOrderProjection(event.OrderId)

	emit.EmitOrderPlaced(event)

	for _, article := range event.PlaceEvent.Articles {
		go emit.EmitArticleValidation(emit.ArticleValidationData{
			ReferenceId: event.OrderId,
			ArticleId:   article.ArticleId,
		})
	}

	return event, err
}
