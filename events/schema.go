package events

import (
	"time"

	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type PaymentMethod string

const (
	Cash   PaymentMethod = "CASH"
	Credit PaymentMethod = "CREDIT"
	Debit  PaymentMethod = "DEBIT"
)

type EventType string

const (
	Place      EventType = "place_order"
	Validation EventType = "aticle_validation"
	Payment    EventType = "payment"
)

// Estuctura basica de del evento
type Event struct {
	ID         string
	OrderId    string    `validate:"required,min=1,max=100"`
	Type       EventType `validate:"required"`
	PlaceEvent *PlaceEvent
	Validation *ValidationEvent
	Payment    *PaymentEvent
	Created    time.Time
	Updated    time.Time
}

// ValidateSchema valida la estructura para ser insertada en la db
func (e *Event) ValidateSchema() error {
	return validator.New().Struct(e)
}

type PlaceEvent struct {
	CartId   string    `json:"cartId"`
	UserId   string    `json:"userId" `
	Articles []Article `json:"articles" `
}

type Article struct {
	ArticleId string `json:"articleId" binding:"required,min=1,max=100"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type PaymentEvent struct {
	OrderId string        `json:"orderId" binding:"required"`
	Method  PaymentMethod `json:"metod" binding:"required"`
	Amount  float32       `json:"amount" binding:"required"`
}

type ValidationEvent struct {
	ArticleId   string  `json:"articleId" json:"articleId"`
	ReferenceId string  `json:"referenceId" json:"referenceId"`
	IsValid     bool    `json:"isValid" json:"valid"`
	Stock       int     `json:"stock" json:"stock"`
	Price       float32 `json:"price" json:"price"`
}

// NewPlaceEvent Nueva instancia de place event
func newPlaceEvent(
	event *PlaceEvent,
) *Event {
	return &Event{
		ID:         uuid.NewV4().String(),
		OrderId:    uuid.NewV4().String(),
		Type:       Place,
		PlaceEvent: event,
		Created:    time.Now(),
		Updated:    time.Now(),
	}
}

// PaymentEvent Nueva instancia de payment event
func newPaymentEvent(
	paymentEvent *PaymentEvent,
) *Event {
	return &Event{
		ID:      uuid.NewV4().String(),
		OrderId: paymentEvent.OrderId,
		Type:    Payment,
		Payment: paymentEvent,
		Created: time.Now(),
		Updated: time.Now(),
	}
}

// ValidationEvent Nueva instancia de validation event
func newValidationEvent(
	validationEvent *ValidationEvent,
) *Event {
	return &Event{
		ID:         uuid.NewV4().String(),
		OrderId:    validationEvent.ReferenceId,
		Type:       Validation,
		Validation: validationEvent,
		Created:    time.Now(),
		Updated:    time.Now(),
	}
}
