package auth

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRequest struct {
	Username string `json:"uname"`
	Password string `json:"upwd"`
	Salt     string `json:"salt"`
}

type loginResponse struct {
	Uid       int64  `json:"uid"`
	Authority int64  `json:"authority"`
	Token     string `json:"token"`
}

// LoginNormal godoc
// @Summary 使用用户名密码的常规登陆方式
// @Schemes
// @Tags example
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Param loginRequest body loginRequest true "loginRequest 认证信息"
// @Success 200 {object} app.Response
// @Router /api/loginRequest/login_normal [post]
func LoginNormal(c *gin.Context) {
	appG := app.Gin{C: c}

	var authInstance loginRequest
	if err := c.BindJSON(&authInstance); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{
		Username: authInstance.Username,
	}

	if err := userService.FillUidByUserName(); err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_LOGIN_WRONG_UNAME_PWD, nil)
		return
	}

	if !userService.CheckPassword(authInstance.Password, authInstance.Salt) {
		appG.Response(http.StatusOK, e.ERROR_AUTH_LOGIN_WRONG_UNAME_PWD, nil)
		return
	}

	if err := userService.FillAuthorityByUid(); err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
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

	appG.Response(http.StatusOK, e.SUCCESS, loginResponse{
		Uid:       userService.UID,
		Authority: userService.Authority,
		Token:     token,
	})
}
