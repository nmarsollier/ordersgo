package services

import (
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/projections"
)

func ProcessArticleData(data *events.ValidationEvent) (*events.Event, error) {
	event, err := events.SaveArticleExist(data)
	if err != nil {
		return nil, err
	}

	go projections.Update(event.OrderId)

	return event, err
}
