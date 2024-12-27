package services

import (
	"github.com/nmarsollier/ordersgo/internal/engine/log"
	"github.com/nmarsollier/ordersgo/internal/events"
	"github.com/nmarsollier/ordersgo/internal/projections"
	"github.com/nmarsollier/ordersgo/internal/rabbit/emit"
)

type Service interface {
	ProcessArticleData(data *events.ValidationEvent) (*events.Event, error)
	PocessPlaceOrder(data *events.PlacedOrderData) (*events.Event, error)
	ProcessSavePayment(data *events.PaymentEvent) (*events.Event, error)
}

func NewService(log log.LogRusEntry, events events.EventService, projections projections.ProjectionsService, emit emit.RabbitEmit) Service {
	return &service{
		log:         log,
		events:      events,
		projections: projections,
		emit:        emit,
	}
}

type service struct {
	log         log.LogRusEntry
	events      events.EventService
	projections projections.ProjectionsService
	emit        emit.RabbitEmit
}

func (s *service) ProcessArticleData(data *events.ValidationEvent) (*events.Event, error) {
	event, err := s.events.SaveArticleExist(data)
	if err != nil {
		return nil, err
	}

	go s.projections.Update(event.OrderId)

	return event, err
}

func (s *service) PocessPlaceOrder(data *events.PlacedOrderData) (*events.Event, error) {
	event, err := s.events.SavePlaceOrder(data)
	if err != nil {
		return nil, err
	}

	go s.projections.Update(event.OrderId)

	go s.emit.EmitOrderPlaced(event)

	for _, article := range event.PlaceEvent.Articles {
		go s.emit.EmitArticleValidation(emit.ArticleValidationData{
			ReferenceId: event.OrderId,
			ArticleId:   article.ArticleId,
		})
	}

	return event, err
}

func (s *service) ProcessSavePayment(data *events.PaymentEvent) (*events.Event, error) {
	event, err := s.events.SavePayment(data)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	go s.projections.Update(event.OrderId)

	return event, err
}
