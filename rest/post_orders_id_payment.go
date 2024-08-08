package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
	"github.com/nmarsollier/ordersgo/services"
)

// Agrega un Pago
//
//	@Summary		Agrega un Pago
//	@Description	Agrega un Pago
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			orderId			path		string					true	"ID de orden"
//	@Param			Authorization	header		string					true	"bearer {token}"
//	@Param			body			body		events.PaymentEvent		true	"Informacion del pago"
//	@Success		200				{object}	order_proj.Order		"Ordenes"
//	@Failure		400				{object}	errors.ErrValidation	"Bad Request"
//	@Failure		401				{object}	errors.ErrCustom		"Unauthorized"
//	@Failure		404				{object}	errors.ErrCustom		"Not Found"
//	@Failure		500				{object}	errors.ErrCustom		"Internal Server Error"
//
//	@Router			/v1/orders/:orderId/payment [post]
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

	event, err := services.ProcessSavePayment(&body)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, event)
}
