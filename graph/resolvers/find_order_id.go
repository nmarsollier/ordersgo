package resolvers

import (
	"context"

	"github.com/nmarsollier/ordersgo/graph/model"
	"github.com/nmarsollier/ordersgo/graph/tools"
	"github.com/nmarsollier/ordersgo/projections/order"
)

func FindByOrderId(ctx context.Context, id string) (*model.Order, error) {
	env := tools.GqlCtx(ctx)
	order, err := order.FindByOrderId(id, env...)
	if err != nil {
		return nil, err
	}

	return mapOrderToModel(order), nil
}
