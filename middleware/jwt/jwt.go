package jwt

import (
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/auth_service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

//// JWT is jwt middleware
//func JWT() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var code int
//		var data interface{}
//
//		code = e.SUCCESS
//		token := c.Query("token")
//		if token == "" {
//			code = e.INVALID_PARAMS
//		} else {
//			_, err := util.ParseToken(token)
//			if err != nil {
//				switch err.(*jwt.ValidationError).Errors {
//				case jwt.ValidationErrorExpired:
//					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
//				default:
//					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
//				}
//			}
//		}
//
//		if code != e.SUCCESS {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"code": code,
//				"msg":  e.GetMsg(code),
//				"data": data,
//			})
//
//			c.Abort()
//			return
//		}
//
//		c.Next()
//	}
//}

func getClaims(c *gin.Context) (*util.Claims, int) {
	token := c.Query("token")

	if token == "" {
		return nil, e.INVALID_PARAMS
	}

	code := e.SUCCESS
	claims, err := util.ParseToken(token)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		default:
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
	}

	return claims, code
}

// AuthLogin whether the user has logged in
func AuthLogin(c *gin.Context) {
	_, code := getClaims(c)

	if code != e.SUCCESS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})

		c.Abort()
		return
	}

	c.Next()
}

func AuthRoot(c *gin.Context) {
	claims, code := getClaims(c)

	authService := auth_service.Auth{Uid: claims.Uid}
	if code == e.SUCCESS && !authService.IsRoot() {
		code = e.ERROR_AUTH_PERMISSION_DENIED
	}

	if code != e.SUCCESS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})

		c.Abort()
		return
	}

	c.Next()
}

func AuthAdmin(c *gin.Context) {
	claims, code := getClaims(c)

	authService := auth_service.Auth{Uid: claims.Uid}
	if code == e.SUCCESS && !authService.IsAdmin() {
		code = e.ERROR_AUTH_PERMISSION_DENIED
	}

	if code != e.SUCCESS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})

		c.Abort()
		return
	}

	c.Next()
}
