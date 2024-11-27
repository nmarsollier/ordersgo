package resolvers

import (
	"context"

	"github.com/nmarsollier/ordersgo/graph/model"
	"github.com/nmarsollier/ordersgo/graph/tools"
	"github.com/nmarsollier/ordersgo/projections/order"
)

func GetOrders(ctx context.Context) ([]*model.OrderSummary, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	env := tools.GqlDeps(ctx)
	e, err := order.FindByUserId(user.ID, env...)
	if err != nil {
		return nil, err
	}

	orders := []*model.OrderSummary{}
	for _, o := range e {
		orders = append(orders, &model.OrderSummary{
			ID:           o.OrderId,
			Status:       model.OrderStatus(o.Status),
			CartID:       o.CartId,
			TotalPrice:   float64(o.TotalPrice()),
			TotalPayment: float64(o.TotalPayment()),
			Articles:     len(o.Articles),
		})
	}

	return orders, nil
}
