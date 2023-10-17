package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/order_proj"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

/**
 * @api {get} /v1/orders/:orderId Buscar Orden
 * @apiName Buscar Orden
 * @apiGroup Ordenes
 *
 * @apiDescription Busca una order del usuario logueado, por su id.
 *
 * @apiUse AuthHeader
 *
 * @apiSuccessExample {json} Respuesta
 *   HTTP/1.1 200 OK
 *   {
 *      "id": "{orderID}",
 *      "status": "{Status}",
 *      "cartId": "{cartId}",
 *      "updated": "{updated date}",
 *      "created": "{created date}",
 *      "articles": [
 *         {
 *             "id": "{articleId}",
 *             "quantity": {quantity},
 *             "validated": true|false,
 *             "valid": true|false
 *         }, ...
 *     ]
 *   }
 *
 * @apiUse Errors
 */
func init() {
	engine.Router().GET(
		"/v1/orders/:orderId",
		middlewares.ValidateAuthentication,
		getOrderById,
	)
}

func getOrderById(c *gin.Context) {
	orderId := c.Param("orderId")

	order, err := order_proj.FindById(orderId)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, order)
}
