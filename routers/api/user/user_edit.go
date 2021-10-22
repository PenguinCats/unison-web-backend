package user

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/auth_service"
	"github.com/PenguinCats/unison-web-backend/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type editUserRequest struct {
	Uid                 int64   `json:"uid"`
	Name                string  `json:"name"`
	Authority           int64   `json:"authority"`
	SeuID               string  `json:"seu_id"`
	PermissionGroupList []int64 `json:"permission_group_list"`
}

func EditUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var req editUserRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	authServiceTarget := auth_service.Auth{
		Uid: req.Uid,
	}

	uid, code := util.ParseUidFromContext(appG.C)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}
	authService := auth_service.Auth{
		Uid: uid,
	}
	authUid, code := authService.GetAuthority()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	if (authServiceTarget.IsAdmin() && uid != req.Uid) || (uid == req.Uid && req.Authority != authUid) {
		appG.Response(http.StatusOK, e.ERROR_AUTH_PERMISSION_DENIED, nil)
		return
	}

	userService := user_service.User{
		UID:             req.Uid,
		Name:            req.Name,
		Authority:       req.Authority,
		SeuID:           req.SeuID,
		PermissionGroup: req.PermissionGroupList,
	}
	code = userService.EditUserUnion()
	appG.Response(http.StatusOK, code, nil)
}
