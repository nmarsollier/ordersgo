package status

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Estuctura basica de del evento
type OrderStatus struct {
	ID      string `dynamodbav:"id,omitempty" json:"id"`
	OrderId string `dynamodbav:"orderId"`
	UserId  string `dynamodbav:"userId"`

	Placed           bool `dynamodbav:"placed"`
	PartialValidated bool `dynamodbav:"partialValidated"`
	Validated        bool `dynamodbav:"validated"`
	PartialPayment   bool `dynamodbav:"partialPayment"`
	PaymentCompleted bool `dynamodbav:"paymentCompleted"`

	Created time.Time `dynamodbav:"created"`
	Updated time.Time `dynamodbav:"updated"`
}

// ValidateSchema valida la estructura para ser insertada en la db
func (e *OrderStatus) ValidateSchema() error {
	return validator.New().Struct(e)
}
