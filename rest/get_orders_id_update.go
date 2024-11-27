package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/projections"
	"github.com/nmarsollier/ordersgo/rest/server"
)

//	@Summary		Actualiza la proyeccion
//	@Description	Actualiza las proyecciones en caso que hayamos roto algo.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"Bearer {token}"
//	@Param			orderId			path	string	true	"ID de orden"
//	@Success		200				"No Content"
//	@Router			/v1/orders/:orderId/update [get]
//
// Updates the Porjections
func init() {
	server.Router().GET(
		"/v1/orders/:orderId/update",
		server.ValidateAuthentication,
		updateOrderById,
	)
}

func updateOrderById(c *gin.Context) {
	orderId := c.Param("orderId")

	deps := server.GinDeps(c)
	go projections.Update(orderId, deps...)

	c.JSON(200, "")
}
