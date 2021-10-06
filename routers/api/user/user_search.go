package user

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/user_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type userSearchRequest struct {
	Query string `json:"query"`
}

type userSearchResponseItem struct {
	Value     int64  `json:"value"`
	UID       int64  `json:"uid"`
	Name      string `json:"name"`
	Authority int64  `json:"authority"`
	SeuID     string `json:"seu_id"`
}

type userSearchResponse struct {
	Users []userSearchResponseItem `json:"users"`
}

func GetUserSearch(c *gin.Context) {
	appG := app.Gin{C: c}
	var req userSearchRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{
		Query: req.Query,
	}

	users, err := userService.GetUserProfileByQueryString()
	if users == nil || err != nil && err != gorm.ErrRecordNotFound {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	res := userSearchResponse{Users: []userSearchResponseItem{}}
	for _, item := range *users {
		res.Users = append(res.Users, userSearchResponseItem{
			Value:     item.UID,
			UID:       item.UID,
			Name:      item.Name,
			Authority: item.Authority,
			SeuID:     item.SeuID,
		})
	}
	appG.Response(http.StatusOK, e.SUCCESS, res)
}
