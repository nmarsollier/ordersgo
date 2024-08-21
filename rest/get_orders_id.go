package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/projections/order"
	"github.com/nmarsollier/ordersgo/rest/server"
)

//	@Summary		Buscar Orden
//	@Description	Busca una order del usuario logueado, por su id.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			orderId			path		string				true	"ID de orden"
//	@Param			Authorization	header		string				true	"bearer {token}"
//	@Success		200				{object}	order.Order			"Ordenes"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	server.ErrorData	"Unauthorized"
//	@Failure		404				{object}	server.ErrorData	"Not Found"
//	@Failure		500				{object}	server.ErrorData	"Internal Server Error"
//	@Router			/v1/orders/:orderId [get]
//
// Buscar Orden
func init() {
	server.Router().GET(
		"/v1/orders/:orderId",
		server.ValidateAuthentication,
		getOrderById,
	)
}

func getOrderById(c *gin.Context) {
	orderId := c.Param("orderId")

	ctx := server.GinCtx(c)
	order, err := order.FindByOrderId(orderId, ctx...)
	if err != nil {
		server.AbortWithError(c, err)
		return
	}

	c.JSON(200, order)
}
