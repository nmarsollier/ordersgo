package rest

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/order_proj"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
	"github.com/nmarsollier/ordersgo/security"
)

// Ordenes de Usuario
//
//	@Summary		Ordenes de Usuario
//	@Description	Busca todas las ordenes del usuario logueado.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"bearer {token}"
//	@Success		200				{array}		OrderListData		"Ordenes"
//	@Failure		400				{object}	errors.ErrValidation	"Bad Request"
//	@Failure		401				{object}	errors.ErrCustom		"Unauthorized"
//	@Failure		404				{object}	errors.ErrCustom		"Not Found"
//	@Failure		500				{object}	errors.ErrCustom		"Internal Server Error"
//
//	@Router			/v1/orders [get]
func init() {
	engine.Router().GET(
		"/v1/orders",
		middlewares.ValidateAuthentication,
		getOrders,
	)
}

func getOrders(c *gin.Context) {
	tokenString, err := middlewares.GetHeaderToken(c)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	user, err := security.Validate(tokenString)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	e, err := order_proj.FindByUserId(user.ID)
	if err != nil {
		middlewares.AbortWithError(c, err)
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

func totalPayment(order *order_proj.Order) float32 {
	var result float32 = 0
	for _, o := range order.Payments {
		result += float32(o.Amount)
	}
	return result
}

func totalPrice(order *order_proj.Order) float32 {
	var result float32 = 0
	for _, o := range order.Articles {
		result += float32(o.UnitaryPrice) * float32(o.Quantity)
	}
	return result
}

type OrderListData struct {
	Id           string                 `json:"id"`
	Status       order_proj.OrderStatus `json:"status"`
	CartId       string                 `json:"cartId"`
	TotalPrice   float32                `json:"totalPrice"`
	TotalPayment float32                `json:"totalPayment"`
	Updated      time.Time              `json:"updated"`
	Created      time.Time              `json:"created"`
	Articles     int                    `json:"articles"`
}
