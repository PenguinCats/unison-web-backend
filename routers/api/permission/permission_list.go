package permission

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/auth_service"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type permissionGroupItem struct {
	Name  string `json:"name"`
	ID    int64  `json:"id"`
	Value int64  `json:"value"`
}
type getPermissionGroupListResponse struct {
	Groups []permissionGroupItem `json:"groups"`
}

func GetPermissionList(c *gin.Context) {
	appG := app.Gin{C: c}

	uid, code := util.ParseUidFromContext(appG.C)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	authService := auth_service.Auth{
		Uid: uid,
	}

	if !authService.IsAdmin() {
		appG.Response(http.StatusOK, e.ERROR_AUTH_PERMISSION_DENIED, nil)
		return
	}

	permissionGroupService := permission_group_service.PermissionGroupService{}

	groups, code := permissionGroupService.GetPermissionGroupList()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	response := getPermissionGroupListResponse{Groups: []permissionGroupItem{}}
	for _, v := range *groups {
		response.Groups = append(response.Groups, permissionGroupItem{
			Name:  v.Name,
			ID:    v.GroupID,
			Value: v.GroupID,
		})
	}

	appG.Response(http.StatusOK, code, response)
	return
}
