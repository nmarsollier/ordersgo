package status

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Estuctura basica de del evento
type OrderStatus struct {
	ID      string `json:"id"`
	OrderId string `json:"orderId"`
	UserId  string `json:"userId"`

	Placed           bool `json:"placed"`
	PartialValidated bool `json:"partialValidated"`
	Validated        bool `json:"validated"`
	PartialPayment   bool `json:"partialPayment"`
	PaymentCompleted bool `json:"paymentCompleted"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// ValidateSchema valida la estructura para ser insertada en la db
func (e *OrderStatus) ValidateSchema() error {
	return validator.New().Struct(e)
}
