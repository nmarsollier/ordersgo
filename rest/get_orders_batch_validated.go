package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

// Batch Validated
//
//	@Summary		Batch Validated
//	@Description	Ejecuta un proceso batch para ordenes en estado VALIDATED.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"bearer {token}"
//	@Success		200				"No Content"
//	@Router			/v1/orders_batch/validated [get]
func init() {
	engine.Router().GET(
		"/v1/orders_batch/validated",
		middlewares.ValidateAuthentication,
		batchValidated,
	)
}

func batchValidated(c *gin.Context) {

	c.JSON(200, "")
}
