package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
)

/**
 * @api {post} /v1/orders/:orderId/payment Agregar Pago
 * @apiName Agrega un Pago
 * @apiGroup Pagos
 *
 * @apiUse AuthHeader
 *
 * @apiExample {json} Body
 *   {
 *       "paymentMethod": "CASH | CREDIT | DEBIT",
 *       "amount": "{amount}"
 *   }
 *
 * @apiSuccessExample {json} Respuesta
 *   HTTP/1.1 200 OK
 *
 * @apiUse Errors
 */
func init() {
	engine.Router().POST(
		"/v1/orders/:orderId/payment",
		middlewares.ValidateAuthentication,
		savePayment,
	)
}

func savePayment(c *gin.Context) {
	body := c.MustGet("data").(events.PaymentEvent)
	if err := c.ShouldBindJSON(&body); err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	event, err := events.SavePayment(&body)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, event)
}
