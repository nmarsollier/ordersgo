package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

/**
 * @api {get} /v1/orders_batch/placed Batch Placed
 * @apiName Batch Placed
 * @apiGroup Ordenes
 *
 * @apiDescription Ejecuta un proceso batch que chequea ordenes en estado PLACED.
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
		"/v1/orders_batch/placed",
		middlewares.ValidateAuthentication,
		batchPlaced,
	)
}

func batchPlaced(c *gin.Context) {

	c.JSON(200, "")
}
