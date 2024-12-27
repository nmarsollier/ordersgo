package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/internal/rest/engine"
)

//	@Summary		Buscar Orden
//	@Description	Busca una order del usuario logueado, por su id.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			orderId			path		string				true	"ID de orden"
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Success		200				{object}	order.Order			"Ordenes"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	engine.ErrorData	"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData	"Not Found"
//	@Failure		500				{object}	engine.ErrorData	"Internal Server Error"
//	@Router			/orders/:orderId [get]
//
// Buscar Orden
func init() {
	engine.Router().GET(
		"/orders/:orderId",
		engine.ValidateAuthentication,
		getOrderById,
	)
}

func getOrderById(c *gin.Context) {
	orderId := c.Param("orderId")

	deps := engine.GinDi(c)
	order, err := deps.OrderService().FindByOrderId(orderId)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, order)
}
