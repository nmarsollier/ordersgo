package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/order_projection"
	"github.com/nmarsollier/ordersgo/rest/engine"
)

//	@Summary		Actualiza la proyeccion
//	@Description	Actualiza las proyecciones en caso que hayamos roto algo.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"bearer {token}"
//	@Param			orderId			path	string	true	"ID de orden"
//	@Success		200				"No Content"
//	@Router			/v1/orders/:orderId/update [get]
//
// Updates the Porjections
func init() {
	engine.Router().GET(
		"/v1/orders/:orderId/update",
		engine.ValidateAuthentication,
		updateOrderById,
	)
}

func updateOrderById(c *gin.Context) {
	orderId := c.Param("orderId")

	order_projection.UpdateOrderProjection(orderId)

	c.JSON(200, "")
}
