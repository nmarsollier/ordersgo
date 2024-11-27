package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/ordersgo/security"
	"github.com/nmarsollier/ordersgo/tools/errs"
	"github.com/nmarsollier/ordersgo/tools/log"
)

// ValidateAuthentication validate gets and check variable body to create new variable
// puts model.Variable in context as body if everything is correct
func ValidateAuthentication(c *gin.Context) {
	user, err := validateToken(c)

	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	deps := GinDeps(c)
	c.Set("logger", log.Get(deps...).WithField(log.LOG_FIELD_USER_ID, user.ID))
}

// get token from Authorization header
func HeaderToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(strings.ToUpper(tokenString), "BEARER ") != 0 {
		deps := GinDeps(c)
		log.Get(deps...).Error(errs.Unauthorized)
		return "", errs.Unauthorized
	}
	return tokenString[7:], nil
}

func validateToken(c *gin.Context) (*security.User, error) {
	tokenString, err := HeaderToken(c)
	if err != nil {
		deps := GinDeps(c)

		log.Get(deps...).Error(err)
		return nil, errs.Unauthorized
	}

	user, err := security.Validate(tokenString)
	if err != nil {
		deps := GinDeps(c)
		log.Get(deps...).Error(err)
		return nil, errs.Unauthorized
	}

	return user, nil
}
