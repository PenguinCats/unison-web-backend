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

type userProfileRequest struct {
	Uid int64 `json:"uid"`
}

type userProfileResponse struct {
	Username  string `json:"username"`
	Uid       int64  `json:"uid"`
	Name      string `json:"name"`
	Authority int64  `json:"authority"`
	SeuID     string `json:"seu_id"`
}

func GetUserProfile(c *gin.Context) {
	appG := app.Gin{C: c}
	var req userProfileRequest

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
	if !authService.IsSelfOrAdmin(req.Uid) {
		appG.Response(http.StatusOK, e.ERROR_AUTH_PERMISSION_DENIED, nil)
		return
	}

	userService := user_service.User{UID: req.Uid}
	err := userService.GetUserProfileByUid()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, userProfileResponse{
		Uid:       userService.UID,
		Name:      userService.Name,
		Username:  userService.Username,
		Authority: userService.Authority,
		SeuID:     userService.SeuID,
	})
}
