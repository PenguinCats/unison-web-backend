package permission

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updatePermissionGroupNameRequest struct {
	Gid  int64  `json:"gid"`
	Name string `json:"name"`
}

func UpdatePermissionGroupName(c *gin.Context) {
	appG := app.Gin{C: c}
	var req updatePermissionGroupNameRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	pgs := permission_group_service.PermissionGroupService{
		GroupID: req.Gid,
		Name:    req.Name,
	}
	code := pgs.EditGroupName()
	appG.Response(http.StatusOK, code, nil)
}
