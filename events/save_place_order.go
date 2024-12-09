package events

import (
	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/ordersgo/tools/log"
)

// SavePlaceOrder saves the event for place order
func SavePlaceOrder(data *PlacedOrderData, deps ...interface{}) (*Event, error) {
	if e, _ := findPlaceByCartId(data.CartId, deps...); e != nil {
		log.Get(deps...).Error("Place already exist")
		return e, nil
	}

	if err := validator.New().Struct(data); err != nil {
		log.Get(deps...).Error("Invalid NewPlaceData Data", err)
		return nil, err
	}

	event := placeOrderToEvent(data)
	event, err := insert(event, deps...)

	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return event, nil
}

type PlacedOrderData struct {
	CartId   string                  `json:"cartId" binding:"required,min=1,max=100"`
	UserId   string                  `json:"userId" binding:"required,min=1,max=100"`
	Articles []PlacePrderArticleData `json:"articles" binding:"required,gt=0"`
}

type PlacePrderArticleData struct {
	Id       string `json:"id" binding:"required,min=1,max=100"`
	Quantity int    `json:"quantity" binding:"required,min=1"`
}

func placeOrderToEvent(event *PlacedOrderData) *Event {
	articles := make([]Article, len(event.Articles))
	for index, item := range event.Articles {
		articles[index] = Article{
			ArticleId: item.Id,
			Quantity:  item.Quantity,
		}
	}

	return newPlaceEvent(&PlaceEvent{
		CartId:   event.CartId,
		UserId:   event.UserId,
		Articles: articles,
	})
}
