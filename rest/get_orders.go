package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/order_proj"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
	"github.com/nmarsollier/ordersgo/security"
)

/**
 * @api {get} /v1/orders Ordenes de Usuario
 * @apiName Ordenes de Usuario
 * @apiGroup Ordenes
 *
 * @apiDescription Busca todas las ordenes del usuario logueado.
 *
 * @apiUse AuthHeader
 *
 *  @apiSuccessExample {json} Respuesta
 *   HTTP/1.1 200 OK
 *   [{
 *      "id": "{orderID}",
 *      "status": "{Status}",
 *      "cartId": "{cartId}",
 *      "updated": "{updated date}",
 *      "created": "{created date}",
 *      "totalPrice": {price}
 *      "articles": {count}
 *   }, ...
 *   ]
 * @apiUse Errors
 */
func init() {
	engine.Router().GET(
		"/v1/orders",
		middlewares.ValidateAuthentication,
		getOrders,
	)
}

func getOrders(c *gin.Context) {
	tokenString, err := middlewares.GetHeaderToken(c)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	user, err := security.Validate(tokenString)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	e, err := order_proj.FindByUserId(user.ID)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, e)
}
