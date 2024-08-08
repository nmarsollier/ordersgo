package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/order_proj"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

// Buscar Orden
//
//	@Summary		Buscar Orden
//	@Description	Busca una order del usuario logueado, por su id.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			orderId			path		string					true	"ID de orden"
//	@Param			Authorization	header		string					true	"bearer {token}"
//	@Success		200				{object}	order_proj.Order		"Ordenes"
//	@Failure		400				{object}	errors.ErrValidation	"Bad Request"
//	@Failure		401				{object}	errors.ErrCustom		"Unauthorized"
//	@Failure		404				{object}	errors.ErrCustom		"Not Found"
//	@Failure		500				{object}	errors.ErrCustom		"Internal Server Error"
//
//	@Router			/v1/orders/:orderId [get]
func init() {
	engine.Router().GET(
		"/v1/orders/:orderId",
		middlewares.ValidateAuthentication,
		getOrderById,
	)
}

func getOrderById(c *gin.Context) {
	orderId := c.Param("orderId")

	order, err := order_proj.FindByOrderId(orderId)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, order)
}
