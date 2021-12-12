package permission

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type bindPermissionGroupHostRequest struct {
	Gid  int64   `json:"gid"`
	Hids []int64 `json:"hids"`
}

func BindPermissionGroupHost(c *gin.Context) {
	appG := app.Gin{C: c}
	var req bindPermissionGroupHostRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	pghs := permission_group_service.PermissionGroupHostService{
		GroupID: req.Gid,
		HIDs:    req.Hids,
	}
	code := pghs.BindHosts()
	appG.Response(http.StatusOK, code, nil)
}
