package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

// Batch Placed
//
//	@Summary		Batch Placed
//	@Description	Ejecuta un proceso batch para ordenes en estado PLACED.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"bearer {token}"
//	@Success		200				"No Content"
//	@Router			/v1/orders_batch/placed [get]
func init() {
	engine.Router().GET(
		"/v1/orders_batch/placed",
		middlewares.ValidateAuthentication,
		batchPlaced,
	)
}

func batchPlaced(c *gin.Context) {

	c.JSON(200, "")
}
