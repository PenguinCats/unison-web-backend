package jwt

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthLogin whether the user has logged in
func AuthLogin(c *gin.Context) {
	appG := app.Gin{C: c}

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		c.Abort()
		return
	}

	_, code := util.CheckClaims(token)

	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})

		c.Abort()
		return
	}

	c.Next()
}

func AuthRoot(c *gin.Context) {
	appG := app.Gin{C: c}

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		c.Abort()
		return
	}

	claims, code := util.CheckClaims(token)

	authService := auth_service.Auth{Uid: claims.Uid}
	if code == e.SUCCESS && !authService.IsRoot() {
		code = e.ERROR_AUTH_PERMISSION_DENIED
	}

	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})

		c.Abort()
		return
	}

	c.Next()
}

func AuthAdmin(c *gin.Context) {
	appG := app.Gin{C: c}

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		c.Abort()
		return
	}

	claims, code := util.CheckClaims(token)
	if code == e.SUCCESS {
		authService := auth_service.Auth{Uid: claims.Uid}
		if !authService.IsAdmin() {
			code = e.ERROR_AUTH_PERMISSION_DENIED
		}
	}

	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})

		c.Abort()
		return
	}

	c.Next()
}
