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

type userPasswordEditRequest struct {
	Uid      int64  `json:"uid"`
	Password string `json:"password"`
}

func EditUserPassword(c *gin.Context) {
	appG := app.Gin{C: c}
	var req userPasswordEditRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	uid, code := util.ParseUidFromContext(appG.C)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	if uid != req.Uid {
		authServiceTarget := auth_service.Auth{
			Uid: req.Uid,
		}
		authServiceContext := auth_service.Auth{
			Uid: uid,
		}
		authTarget, code := authServiceTarget.GetAuthority()
		if code != e.SUCCESS {
			appG.Response(http.StatusOK, code, nil)
			return
		}
		authContext, code := authServiceContext.GetAuthority()
		if code != e.SUCCESS {
			appG.Response(http.StatusOK, code, nil)
			return
		}
		if authTarget <= authContext {
			code = e.ERROR_AUTH_PERMISSION_DENIED
			appG.Response(http.StatusOK, code, nil)
			return
		}
	}

	userService := user_service.User{
		UID:      req.Uid,
		Password: req.Password,
	}
	code = userService.EditUserPassword()
	appG.Response(http.StatusOK, code, nil)
}
