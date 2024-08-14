package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/services"
)

//	@Summary		Agrega un Pago
//	@Description	Agrega un Pago
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			orderId			path		string					true	"ID de orden"
//	@Param			Authorization	header		string					true	"bearer {token}"
//	@Param			body			body		events.PaymentEvent		true	"Informacion del pago"
//	@Success		200				{object}	order_projection.Order	"Ordenes"
//	@Failure		400				{object}	apperr.ValidationErr	"Bad Request"
//	@Failure		401				{object}	engine.ErrorData		"Unauthorized"
//	@Failure		404				{object}	engine.ErrorData		"Not Found"
//	@Failure		500				{object}	engine.ErrorData		"Internal Server Error"
//	@Router			/v1/orders/:orderId/payment [post]
//
// Agrega un Pago
func init() {
	engine.Router().POST(
		"/v1/orders/:orderId/payment",
		engine.ValidateAuthentication,
		savePayment,
	)
}

func savePayment(c *gin.Context) {
	body := events.PaymentEvent{}
	if err := c.ShouldBindJSON(&body); err != nil {
		engine.AbortWithError(c, err)
		return
	}

	event, err := services.ProcessSavePayment(&body)
	if err != nil {
		engine.AbortWithError(c, err)
		return
	}

	c.JSON(200, event)
}
