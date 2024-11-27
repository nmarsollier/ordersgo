package resolvers

import (
	"context"

	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/graph/model"
	"github.com/nmarsollier/ordersgo/graph/tools"
	"github.com/nmarsollier/ordersgo/services"
)

func CreatePayment(ctx context.Context, orderID string, payment *model.PaymentEventInput) (bool, error) {
	_, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlDeps(ctx)

	_, err = services.ProcessSavePayment(&events.PaymentEvent{
		OrderId: orderID,
		Method:  events.PaymentMethod(payment.Method),
		Amount:  float32(payment.Amount),
	}, env...)
	if err != nil {
		return false, err
	}

	return true, nil
}
