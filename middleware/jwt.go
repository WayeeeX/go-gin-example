package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

// JWT is jwt middleware
func AuthJWT(needAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Request.Header.Get("music-access-token")
		if token == "" {
			code = e.ERROR_AUTH_TOKEN_EMPTY
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			} else if claims.Role == 0 && needAdmin {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				c.Set("userID", claims.UserID)
				c.Set("userRole", claims.Role)
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
