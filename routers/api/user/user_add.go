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

type userAddRequest struct {
	UserName            string  `json:"username"`
	Password            string  `json:"password"`
	Name                string  `json:"name"`
	Authority           int64   `json:"authority"`
	SeuId               string  `json:"seu_id"`
	PermissionGroupList []int64 `json:"permission_group_list"`
}

func AddUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var req userAddRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if req.Authority <= 1 || req.Authority > 3 {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	uid, code := util.ParseUidFromContext(appG.C)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}
	authService := auth_service.Auth{
		Uid: uid,
	}

	if req.Authority == 2 && !authService.IsRoot() {
		appG.Response(http.StatusOK, e.ERROR_AUTH_PERMISSION_DENIED, nil)
		return
	}

	userService := user_service.User{
		Username:        req.UserName,
		Password:        req.Password,
		Name:            req.Name,
		Authority:       req.Authority,
		SeuID:           req.SeuId,
		PermissionGroup: req.PermissionGroupList,
	}
	code = userService.AddUserUnion()
	appG.Response(http.StatusOK, code, nil)
	return
}
