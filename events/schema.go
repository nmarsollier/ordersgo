package events

import (
	"time"

	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID         primitive.ObjectID `bson:"_id"`
	OrderId    string             `bson:"orderId" validate:"required,min=1,max=100"`
	Type       EventType          `bson:"type" validate:"required"`
	PlaceEvent *PlaceEvent        `bson:"placeEvent"`
	Validation *ValidationEvent   `bson:"validation"`
	Payment    *PaymentEvent      `bson:"payment"`
	Created    time.Time          `bson:"created"`
	Updated    time.Time          `bson:"updated"`
}

// ValidateSchema valida la estructura para ser insertada en la db
func (e *Event) ValidateSchema() error {
	return validator.New().Struct(e)
}

type PlaceEvent struct {
	CartId   string    `bson:"cartId"`
	UserId   string    `bson:"userId" `
	Articles []Article `bson:"articles" `
}

type Article struct {
	ArticleId string `json:"articleId" binding:"required,min=1,max=100"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type PaymentEvent struct {
	CartId string        `bson:"cartId"`
	Method PaymentMethod `bson:"metod" `
	Amount float32       `bson:"amount" `
}

type ValidationEvent struct {
	ArticleId string  `bson:"articleId"`
	IsValid   bool    `bson:"isValid" `
	Stock     int     `bson:"stock" `
	Price     float32 `bson:"price" `
}

// NewPlaceEvent Nueva instancia de place event
func newPlaceEvent(
	event *PlaceEvent,
) *Event {
	return &Event{
		ID:         primitive.NewObjectID(),
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
		ID:      primitive.NewObjectID(),
		OrderId: uuid.NewV4().String(),
		Type:    Place,
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
		ID:         primitive.NewObjectID(),
		OrderId:    uuid.NewV4().String(),
		Type:       Place,
		Validation: validationEvent,
		Created:    time.Now(),
		Updated:    time.Now(),
	}
}
