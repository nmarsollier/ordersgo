package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/events"
	"github.com/nmarsollier/ordersgo/rest/server"
	"github.com/nmarsollier/ordersgo/services"
)

//	@Summary		Agrega un Pago
//	@Description	Agrega un Pago
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			orderId			path		string				true	"ID de orden"
//	@Param			Authorization	header		string				true	"bearer {token}"
//	@Param			body			body		events.PaymentEvent	true	"Informacion del pago"
//	@Success		200				{object}	order.Order			"Ordenes"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/orders/:orderId/payment [post]
//
// Agrega un Pago
func init() {
	server.Router().POST(
		"/v1/orders/:orderId/payment",
		server.ValidateAuthentication,
		savePayment,
	)
}

func savePayment(c *gin.Context) {
	body := events.PaymentEvent{}
	if err := c.ShouldBindJSON(&body); err != nil {
		server.AbortWithError(c, err)
		return
	}

	ctx := server.GinCtx(c)
	event, err := services.ProcessSavePayment(&body, ctx...)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, event)
}
