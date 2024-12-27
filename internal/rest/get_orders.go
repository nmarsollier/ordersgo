package rest

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/internal/projections/order"
	"github.com/nmarsollier/ordersgo/internal/rest/engine"
)

//	@Summary		Ordenes de Usuario
//	@Description	Busca todas las ordenes del usuario logueado.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Success		200				{array}		OrderListData		"Ordenes"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	engine.ErrorData	"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData	"Not Found"
//	@Failure		500				{object}	engine.ErrorData	"Internal Server Error"
//	@Router			/orders [get]
//
// Ordenes de Usuario
func init() {
	engine.Router().GET(
		"/orders",
		engine.ValidateAuthentication,
		getOrders,
	)
}

func getOrders(c *gin.Context) {
	tokenString, err := engine.HeaderToken(c)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	deps := engine.GinDi(c)
	user, err := deps.SecurityService().Validate(tokenString)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	e, err := deps.OrderService().FindByUserId(user.ID)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	orders := []OrderListData{}
	for _, o := range e {
		orders = append(orders, OrderListData{

			Id:           o.OrderId,
			Status:       o.Status,
			CartId:       o.CartId,
			TotalPrice:   totalPrice(o),
			TotalPayment: totalPayment(o),
			Updated:      o.Updated,
			Created:      o.Created,
			Articles:     len(o.Articles),
		})
	}

	c.JSON(200, orders)
}

func totalPayment(order *order.Order) float32 {
	var result float32 = 0
	for _, o := range order.Payments {
		result += float32(o.Amount)
	}
	return result
}

func totalPrice(order *order.Order) float32 {
	var result float32 = 0
	for _, o := range order.Articles {
		result += float32(o.UnitaryPrice) * float32(o.Quantity)
	}
	return result
}

type OrderListData struct {
	Id           string            `json:"id"`
	Status       order.OrderStatus `json:"status"`
	CartId       string            `json:"cartId"`
	TotalPrice   float32           `json:"totalPrice"`
	TotalPayment float32           `json:"totalPayment"`
	Updated      time.Time         `json:"updated"`
	Created      time.Time         `json:"created"`
	Articles     int               `json:"articles"`
}
