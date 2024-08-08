package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

// Batch Payment Defined
//
//	@Summary		Batch Payment Defined
//	@Description	Ejecuta un proceso batch que chequea ordenes en estado PAYMENT_DEFINED.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"bearer {token}"
//	@Success		200				"No Content"
//	@Router			/v1/orders_batch/payment_defined [get]
func init() {
	engine.Router().GET(
		"/v1/orders_batch/payment_defined",
		middlewares.ValidateAuthentication,
		batchPaymentDefined,
	)
}

func batchPaymentDefined(c *gin.Context) {

	c.JSON(200, "")
}
