package engine

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/security"
	"github.com/nmarsollier/ordersgo/tools/apperr"
)

// ValidateAuthentication validate gets and check variable body to create new variable
// puts model.Variable in context as body if everything is correct
func ValidateAuthentication(c *gin.Context) {
	if err := validateToken(c); err != nil {
		c.Error(err)
		c.Abort()
		return
	}
}

// get token from Authorization header
func HeaderToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "bearer ") != 0 {
		glog.Error(apperr.Unauthorized)
		return "", apperr.Unauthorized
	}
	return tokenString[7:], nil
}

func validateToken(c *gin.Context) error {
	tokenString, err := HeaderToken(c)
	if err != nil {
		glog.Error(err)
		return apperr.Unauthorized
	}

	if _, err = security.Validate(tokenString); err != nil {
		glog.Error(err)
		return apperr.Unauthorized
	}

	return nil
}
