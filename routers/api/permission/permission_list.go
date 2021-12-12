package permission

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type permissionGroupItem struct {
	Name    string `json:"name"`
	ID      int64  `json:"id"`
	Value   int64  `json:"value"`
	HostCnt int64  `json:"host_cnt"`
}
type getPermissionGroupListResponse struct {
	Groups []permissionGroupItem `json:"groups"`
}

func GetPermissionList(c *gin.Context) {
	appG := app.Gin{C: c}

	permissionGroupService := permission_group_service.PermissionGroupService{}

	groups, code := permissionGroupService.GetPermissionGroupList()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	pghs := permission_group_service.PermissionGroupHostService{}
	pghsMap, code := pghs.CountHosts()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	response := getPermissionGroupListResponse{Groups: []permissionGroupItem{}}
	for _, v := range *groups {
		cnt := int64(0)
		if v, ok := pghsMap[v.GroupID]; ok {
			cnt = v
		}
		response.Groups = append(response.Groups, permissionGroupItem{
			Name:    v.Name,
			ID:      v.GroupID,
			Value:   v.GroupID,
			HostCnt: cnt,
		})
	}

	appG.Response(http.StatusOK, code, response)
	return
}
