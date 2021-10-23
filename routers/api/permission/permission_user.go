package permission

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getPermissionGroupOfUserRequest struct {
	Uid int64 `json:"uid"`
}

type getPermissionGroupOfUserResponse struct {
	Gids  []int64  `json:"gids"`
	GName []string `json:"gnames"`
}

func GetPermissionGroupOfUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var req getPermissionGroupOfUserRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	pus := permission_group_service.PermissionUserService{
		UID: req.Uid,
	}
	gids, code := pus.GetGroupIDsByUid()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	var pgs permission_group_service.PermissionGroupService
	groups, code := pgs.GetGroupsByGroupIDs(*gids)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	var res getPermissionGroupOfUserResponse
	if groups != nil {
		for _, item := range *groups {
			res.Gids = append(res.Gids, item.GroupID)
			res.GName = append(res.GName, item.Name)
		}
	}

	appG.Response(http.StatusOK, code, res)
}
