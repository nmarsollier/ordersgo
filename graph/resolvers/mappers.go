package resolvers

import (
	"github.com/nmarsollier/ordersgo/graph/model"
	"github.com/nmarsollier/ordersgo/projections/order"
)

func mapOrderToModel(order *order.Order) *model.Order {
	return &model.Order{
		ID:       order.ID.Hex(),
		OrderID:  order.OrderId,
		Status:   model.OrderStatus(order.Status),
		UserID:   order.UserId,
		CartID:   order.CartId,
		Articles: mapArticlesToModel(order.Articles),
		Payments: mapPaymentsToModel(order.Payments),
	}
}

func mapArticlesToModel(articles []*order.Article) []*model.OrderArticle {
	result := make([]*model.OrderArticle, len(articles))
	for i, a := range articles {
		result[i] = &model.OrderArticle{
			ArticleID:    a.ArticleId,
			Article:      &model.Article{ID: a.ArticleId},
			Quantity:     a.Quantity,
			IsValid:      a.IsValid,
			UnitaryPrice: float64(a.UnitaryPrice),
			IsValidated:  a.IsValidated,
		}
	}
	return result
}

func mapPaymentsToModel(payments []*order.PaymentEvent) []*model.PaymentEvent {
	result := make([]*model.PaymentEvent, len(payments))
	for i, p := range payments {
		result[i] = &model.PaymentEvent{
			Method: model.PaymentMethod(p.Method),
			Amount: float64(p.Amount),
		}
	}
	return result
}
