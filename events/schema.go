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
	ID           string           `dynamodbav:"id,omitempty"`
	OrderId      string           `dynamodbav:"orderId" validate:"required,min=1,max=100"`
	Type         EventType        `dynamodbav:"type" validate:"required"`
	PlaceEvent   *PlaceEvent      `dynamodbav:"placeEvent"`
	Validation   *ValidationEvent `dynamodbav:"validation"`
	Payment      *PaymentEvent    `dynamodbav:"payment"`
	Created      time.Time        `dynamodbav:"created"`
	Updated      time.Time        `dynamodbav:"updated"`
	IdxPlaceCart string           `dynamodbav:"idx_place_cart"`
}

// ValidateSchema valida la estructura para ser insertada en la db
func (e *Event) ValidateSchema() error {
	return validator.New().Struct(e)
}

type PlaceEvent struct {
	CartId   string    `dynamodbav:"cartId"`
	UserId   string    `dynamodbav:"userId" `
	Articles []Article `dynamodbav:"articles" `
}

type Article struct {
	ArticleId string `dynamodbav:"articleId" json:"articleId" binding:"required,min=1,max=100"`
	Quantity  int    `dynamodbav:"quantity" json:"quantity" binding:"required,min=1"`
}

type PaymentEvent struct {
	OrderId string        `dynamodbav:"orderId" binding:"required"`
	Method  PaymentMethod `dynamodbav:"metod" binding:"required"`
	Amount  float32       `dynamodbav:"amount" binding:"required"`
}

type ValidationEvent struct {
	ArticleId   string  `dynamodbav:"articleId" json:"articleId"`
	ReferenceId string  `dynamodbav:"referenceId" json:"referenceId"`
	IsValid     bool    `dynamodbav:"isValid" json:"valid"`
	Stock       int     `dynamodbav:"stock" json:"stock"`
	Price       float32 `dynamodbav:"price" json:"price"`
}

// NewPlaceEvent Nueva instancia de place event
func newPlaceEvent(
	event *PlaceEvent,
) *Event {
	return &Event{
		ID:           uuid.NewV4().String(),
		OrderId:      uuid.NewV4().String(),
		Type:         Place,
		PlaceEvent:   event,
		IdxPlaceCart: event.CartId,
		Created:      time.Now(),
		Updated:      time.Now(),
	}
}

// PaymentEvent Nueva instancia de payment event
func newPaymentEvent(
	paymentEvent *PaymentEvent,
) *Event {
	return &Event{
		ID:           uuid.NewV4().String(),
		OrderId:      paymentEvent.OrderId,
		Type:         Payment,
		Payment:      paymentEvent,
		IdxPlaceCart: "_",
		Created:      time.Now(),
		Updated:      time.Now(),
	}
}

// ValidationEvent Nueva instancia de validation event
func newValidationEvent(
	validationEvent *ValidationEvent,
) *Event {
	return &Event{
		ID:           uuid.NewV4().String(),
		OrderId:      validationEvent.ReferenceId,
		Type:         Validation,
		Validation:   validationEvent,
		IdxPlaceCart: "_",
		Created:      time.Now(),
		Updated:      time.Now(),
	}
}
