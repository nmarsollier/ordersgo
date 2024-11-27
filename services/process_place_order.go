package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections"
	"github.com/nmarsollier/ordersgo/rabbit/emit"
)

func PocessPlaceOrder(data *events.PlacedOrderData, deps ...interface{}) (*events.Event, error) {
	event, err := events.SavePlaceOrder(data, deps...)
	if err != nil {
		return nil, err
	}

	go projections.Update(event.OrderId, deps...)

	go emit.EmitOrderPlaced(event, deps...)

	for _, article := range event.PlaceEvent.Articles {
		go emit.EmitArticleValidation(emit.ArticleValidationData{
			ReferenceId: event.OrderId,
			ArticleId:   article.ArticleId,
		}, deps...)
	}

	return event, err
}
