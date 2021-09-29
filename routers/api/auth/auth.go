package auth

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `json:"uname"`
	Password string `json:"upwd"`
	Salt     string `json:"salt"`
}

// LoginNormal godoc
// @Summary 使用用户名密码的常规登陆方式
// @Schemes
// @Tags example
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Param auth body auth true "auth 认证信息"
// @Success 200 {object} app.Response
// @Router /api/auth/login_normal [post]
func LoginNormal(c *gin.Context) {
	appG := app.Gin{C: c}

	var authInstance auth
	err := c.BindJSON(&authInstance)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{
		Username: authInstance.Username,
	}

	err = userService.FillUidByUserName()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_LOGIN_WRONG_UNAME_PWD, nil)
		return
	}

	if !userService.CheckPassword(authInstance.Password, authInstance.Salt) {
		appG.Response(http.StatusOK, e.ERROR_AUTH_LOGIN_WRONG_UNAME_PWD, nil)
		return
	}

	token, err := util.GenerateToken(userService.UID)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	// 应该重新发起新的请求获取用户信息
	//err = userService.GetUserByUid()
	//if err != nil {
	//	appG.Response(http.StatusOK, e.ERROR, nil)
	//	return
	//}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"uid":   userService.UID,
		"token": token,
	})
}
