package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/internal/rest/engine"
)

//	@Summary		Actualiza la proyeccion
//	@Description	Actualiza las proyecciones en caso que hayamos roto algo.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"Bearer {token}"
//	@Param			orderId			path	string	true	"ID de orden"
//	@Success		200				"No Content"
//	@Router			/orders/:orderId/update [get]
//
// Updates the Porjections
func init() {
	engine.Router().GET(
		"/orders/:orderId/update",
		engine.ValidateAuthentication,
		updateOrderById,
	)
}

func updateOrderById(c *gin.Context) {
	orderId := c.Param("orderId")

	deps := engine.GinDi(c)
	go deps.ProjectionsService().Update(orderId)

	c.JSON(200, "")
}
