package resolvers

import (
	"context"

	"github.com/nmarsollier/ordersgo/graph/model"
	"github.com/nmarsollier/ordersgo/graph/tools"
)

func GetOrder(ctx context.Context, id string) (*model.Order, error) {
	_, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	return FindByOrderId(ctx, id)
}
