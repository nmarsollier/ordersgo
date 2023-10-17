package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

/**
 * @api {get} /v1/orders_batch/validated Batch Validated
 * @apiName Batch Validated
 * @apiGroup Ordenes
 *
 * @apiDescription Ejecuta un proceso batch para ordenes en estado VALIDATED.
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
		"/v1/orders_batch/validated",
		middlewares.ValidateAuthentication,
		batchValidated,
	)
}

func batchValidated(c *gin.Context) {

	c.JSON(200, "")
}
