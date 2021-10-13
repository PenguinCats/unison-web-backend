package user

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/PenguinCats/unison-web-backend/service/user_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type userListRequest struct {
	PageSize   int64 `json:"page_size"`
	PageNumber int64 `json:"page_number"`
}

type userListPermissionGroup struct {
	GroupID   int64  `json:"group_id"`
	GroupName string `json:"group_name"`
}

type userListItem struct {
	Uid             int64                     `json:"uid"`
	UserName        string                    `json:"username"`
	Name            string                    `json:"name"`
	Authority       int64                     `json:"authority"`
	SeuID           string                    `json:"seu_id"`
	PermissionGroup []userListPermissionGroup `json:"permission_group"`
}

type userListResponse struct {
	UserList    []userListItem `json:"user_list"`
	TotalNumber int64          `json:"total_number"`
}

func GetUserList(c *gin.Context) {
	appG := app.Gin{C: c}
	var req userListRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{
		PageSize:   req.PageSize,
		PageNumber: req.PageNumber,
	}

	code := e.SUCCESS
	response := userListResponse{
		UserList:    []userListItem{},
		TotalNumber: 0,
	}

	users, err := userService.GetUserList()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			appG.Response(http.StatusOK, code, response)
			return
		}
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
		return
	}

	var pus permission_group_service.PermissionUserService
	var pgs permission_group_service.PermissionGroupService

	for _, user := range *users {
		pus.UID = user.UID
		gids, code := pus.GetGroupIDsByUid()
		if code != e.SUCCESS {
			appG.Response(http.StatusOK, code, nil)
			return
		}

		var permissionGroup []userListPermissionGroup
		for _, gid := range *gids {
			pgs.GroupID = gid
			gname, code := pgs.GetGroupNameByGroupID()
			if code != e.SUCCESS {
				appG.Response(http.StatusOK, code, nil)
				return
			}
			permissionGroup = append(permissionGroup, userListPermissionGroup{
				GroupID:   gid,
				GroupName: gname,
			})
		}

		response.UserList = append(response.UserList, userListItem{
			Uid:             user.UID,
			UserName:        user.Username,
			Name:            user.Name,
			Authority:       user.Authority,
			SeuID:           user.SeuID,
			PermissionGroup: permissionGroup,
		})
	}

	totalNumber, err := userService.GetTotalUser()
	if err != nil {
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
		return
	}
	response.TotalNumber = totalNumber

	appG.Response(http.StatusOK, code, response)
}
