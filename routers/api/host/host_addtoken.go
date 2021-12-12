package host

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/host_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HostAddTokenResponse struct {
	Token string `json:"token"`
}

func UpdateHostAddToken(c *gin.Context) {
	appG := app.Gin{C: c}

	token, code := host_service.GetHostDaemon().UpdateHostAddToken()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	appG.Response(http.StatusOK, code, HostAddTokenResponse{Token: token})
}

func GetHostAddToken(c *gin.Context) {
	appG := app.Gin{C: c}
	token, code := host_service.GetHostDaemon().GetHostAddToken()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, e.SUCCESS, HostAddTokenResponse{Token: "请刷新密钥并重新获取"})
		return
	}

	appG.Response(http.StatusOK, code, HostAddTokenResponse{Token: token})
}
