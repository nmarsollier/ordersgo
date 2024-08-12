package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/nmarsollier/ordersgo/security"
	"github.com/nmarsollier/ordersgo/tools/apperr"
)

/**
 * @apiDefine AuthHeader
 *
 * @apiExample {String} Header Autorización
 *    Authorization=bearer {token}
 *
 * @apiErrorExample 401 Unauthorized
 *    HTTP/1.1 401 Unauthorized
 */

// ValidateAuthentication validate gets and check variable body to create new variable
// puts model.Variable in context as body if everything is correct
func ValidateAuthentication(c *gin.Context) {
	if err := validateToken(c); err != nil {
		c.Error(err)
		c.Abort()
		return
	}
}

var securityValidate func(token string) (*security.User, error) = security.Validate

func validateToken(c *gin.Context) error {
	tokenString, err := GetHeaderToken(c)
	if err != nil {
		glog.Error(err)
		return apperr.Unauthorized
	}

	if _, err = securityValidate(tokenString); err != nil {
		glog.Error(err)
		return apperr.Unauthorized
	}

	return nil
}

// get token from Authorization header
func GetHeaderToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "bearer ") != 0 {
		glog.Error(apperr.Unauthorized)
		return "", apperr.Unauthorized
	}
	return tokenString[7:], nil
}
