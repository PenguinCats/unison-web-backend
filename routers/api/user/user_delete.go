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

type userDeleteRequest struct {
	Uid int64 `json:"uid"`
}

func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var req userDeleteRequest

	if err := appG.C.BindJSON(&req); err != nil {
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
	if req.Uid == uid {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	if !authService.IsAdmin() {
		appG.Response(http.StatusOK, e.ERROR_AUTH_PERMISSION_DENIED, nil)
		return
	}

	userService := user_service.User{UID: req.Uid}
	code = userService.DeleteUserByUid()
	appG.Response(http.StatusOK, code, nil)
}
