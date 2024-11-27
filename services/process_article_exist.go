package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections"
)

func ProcessArticleData(data *events.ValidationEvent, deps ...interface{}) (*events.Event, error) {
	event, err := events.SaveArticleExist(data, deps...)
	if err != nil {
		return nil, err
	}

	go projections.Update(event.OrderId, deps...)

	return event, err
}
