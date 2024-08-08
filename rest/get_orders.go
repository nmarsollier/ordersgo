package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/order_proj"
	"github.com/nmarsollier/ordersgo/rest/engine"
	"github.com/nmarsollier/ordersgo/rest/middlewares"
	"github.com/nmarsollier/ordersgo/security"
)

// Ordenes de Usuario
//
//	@Summary		Ordenes de Usuario
//	@Description	Busca todas las ordenes del usuario logueado.
//	@Tags			Ordenes
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"bearer {token}"
//	@Success		200				{array}		order_proj.Order		"Ordenes"
//	@Failure		400				{object}	errors.ErrValidation	"Bad Request"
//	@Failure		401				{object}	errors.ErrCustom		"Unauthorized"
//	@Failure		404				{object}	errors.ErrCustom		"Not Found"
//	@Failure		500				{object}	errors.ErrCustom		"Internal Server Error"
//
//	@Router			/v1/orders [get]
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
