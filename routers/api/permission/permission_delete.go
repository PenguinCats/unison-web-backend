package permission

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deletePermissionGroupRequest struct {
	Gid int64 `json:"gid"`
}

func DeletePermissionGroup(c *gin.Context) {
	appG := app.Gin{C: c}
	var req deletePermissionGroupRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	pgs := permission_group_service.PermissionGroupService{
		GroupID: req.Gid,
	}
	code := pgs.DeleteGroup()
	appG.Response(http.StatusOK, code, nil)
}
