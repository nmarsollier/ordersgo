package order

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/ordersgo/events"
)

type OrderStatus string

const (
	Placed          OrderStatus = "placed"
	Invalid         OrderStatus = "invalid"
	Validated       OrderStatus = "validated"
	Payment_Defined OrderStatus = "payment_defined"
)

// Estuctura basica de del evento
type Order struct {
	ID      string      `dynamodbav:"id,omitempty" json:"id"`
	OrderId string      `dynamodbav:"orderId" json:"orderId" validate:"required,min=1,max=100"`
	Status  OrderStatus `dynamodbav:"status" json:"status" validate:"required"`

	UserId   string     `dynamodbav:"userId" json:"userId" validate:"required,min=1,max=100"`
	CartId   string     `dynamodbav:"cartId" json:"cartId" validate:"required,min=1,max=100"`
	Articles []*Article `dynamodbav:"articles"  json:"articles"`

	Payments []*PaymentEvent `dynamodbav:"payments" json:"payments"`

	Created time.Time `dynamodbav:"created" json:"created"`
	Updated time.Time `dynamodbav:"updated" json:"updated"`
}

type Article struct {
	ArticleId    string  `dynamodbav:"articleId" json:"articleId" binding:"required,min=1,max=100"`
	Quantity     int     `dynamodbav:"quantity" json:"quantity" binding:"required,min=1"`
	IsValid      bool    `dynamodbav:"isValid" json:"isValid" `
	UnitaryPrice float32 `dynamodbav:"unitaryPrice" json:"unitaryPrice" `
	IsValidated  bool    `dynamodbav:"isValidated" json:"isValidated" `
}

type PaymentEvent struct {
	Method events.PaymentMethod `dynamodbav:"metod" json:"method"`
	Amount float32              `dynamodbav:"amount" json:"amount"`
}

// ValidateSchema valida la estructura para ser insertada en la db
func (e *Order) ValidateSchema() error {
	return validator.New().Struct(e)
}

func (e *Order) TotalPrice() float32 {
	var result float32
	for _, a := range e.Articles {
		result += a.UnitaryPrice
	}
	return result
}

func (e *Order) TotalPayment() float32 {
	var result float32
	for _, p := range e.Payments {
		result += p.Amount
	}
	return result
}
