package permission

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PermissionGroupHostRequest struct {
	Gid int64 `json:"gid"`
}

type PermissionGroupHostResponse struct {
	Hids []int64 `json:"hids"`
}

func GetHostOfPermissionGroup(c *gin.Context) {
	appG := app.Gin{C: c}
	var req PermissionGroupHostRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	pghs := permission_group_service.PermissionGroupHostService{
		GroupID: req.Gid,
	}
	code := pghs.GetHosts()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	res := PermissionGroupHostResponse{Hids: pghs.HIDs}
	appG.Response(http.StatusOK, code, res)
}
