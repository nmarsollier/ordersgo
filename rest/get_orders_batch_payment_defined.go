package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

/**
 * @api {get} /v1/orders_batch/payment_defined Batch Payment Defined
 * @apiName Batch Payment Defined
 * @apiGroup Ordenes
 *
 * @apiDescription Ejecuta un proceso batch que chequea ordenes en estado PAYMENT_DEFINED.
 *
 * @apiUse AuthHeader
 *
 * @apiSuccessExample {json} Respuesta
 *   HTTP/1.1 200 OK
 *
 *
 * @apiUse Errors
 */
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
