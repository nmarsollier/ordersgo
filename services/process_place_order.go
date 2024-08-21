package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections"
	"github.com/nmarsollier/ordersgo/rabbit/emit"
)

func PocessPlaceOrder(data *events.PlacedOrderData, ctx ...interface{}) (*events.Event, error) {
	event, err := events.SavePlaceOrder(data, ctx...)
	if err != nil {
		return nil, err
	}

	go projections.Update(event.OrderId, ctx...)

	go emit.EmitOrderPlaced(event, ctx...)

	for _, article := range event.PlaceEvent.Articles {
		go emit.EmitArticleValidation(emit.ArticleValidationData{
			ReferenceId: event.OrderId,
			ArticleId:   article.ArticleId,
		}, ctx...)
	}

	return event, err
}
