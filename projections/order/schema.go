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
	ID      string      `json:"id"`
	OrderId string      `json:"orderId" validate:"required,min=1,max=100"`
	Status  OrderStatus `json:"status" validate:"required"`

	UserId   string     `json:"userId" validate:"required,min=1,max=100"`
	CartId   string     `json:"cartId" validate:"required,min=1,max=100"`
	Articles []*Article `json:"articles"`

	Payments []*PaymentEvent `json:"payments"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type Article struct {
	ArticleId    string  `json:"articleId" binding:"required,min=1,max=100"`
	Quantity     int     `json:"quantity" binding:"required,min=1"`
	IsValid      bool    `json:"isValid" `
	UnitaryPrice float32 `json:"unitaryPrice" `
	IsValidated  bool    `json:"isValidated" `
}

type PaymentEvent struct {
	Method events.PaymentMethod `json:"method"`
	Amount float32              `json:"amount"`
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
