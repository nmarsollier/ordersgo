package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections"
)

func ProcessArticleData(data *events.ValidationEvent, ctx ...interface{}) (*events.Event, error) {
	event, err := events.SaveArticleExist(data, ctx...)
	if err != nil {
		return nil, err
	}

	go projections.Update(event.OrderId, ctx...)

	return event, err
}
