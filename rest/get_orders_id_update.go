package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/order_proj"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

// Updates the Porjection
//
//	@Summary		Actualiza la proyeccion
//	@Description	Actualiza las proyecciones en caso que hayamos roto algo.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"bearer {token}"
//	@Param			orderId			path	string	true	"ID de orden"
//	@Success		200				"No Content"
//	@Router			/v1/orders/:orderId/update [get]
func init() {
	engine.Router().GET(
		"/v1/orders/:orderId/update",
		middlewares.ValidateAuthentication,
		updateOrderById,
	)
}

func updateOrderById(c *gin.Context) {
	orderId := c.Param("orderId")

	order_proj.UpdateOrderProjection(orderId)

	c.JSON(200, "")
}
