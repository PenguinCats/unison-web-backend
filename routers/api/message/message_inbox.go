package message

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/message_service"
	"github.com/PenguinCats/unison-web-backend/service/user_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type messageInboxProfileListRequest struct {
	PageSize   int64 `json:"page_size"`
	PageNumber int64 `json:"page_number"`
}

type messageInboxProfile struct {
	ID                int64  `json:"id"`
	Mid               int64  `json:"mid"`
	Title             string `json:"title"`
	Time              string `json:"time"`
	State             int64  `json:"state"`
	FromUserName      string `json:"from_user_name"`
	FromUserAuthority int64  `json:"from_user_authority"`
	FromUserSeuID     string `json:"from_user_seu_id"`
}

type messageInboxProfileListResponse struct {
	MessageProfileList []messageInboxProfile `json:"message_profile_list"`
	TotalNumber        int64                 `json:"total_number"`
}

func GetMessageInboxProfileList(c *gin.Context) {
	appG := app.Gin{C: c}
	var req messageInboxProfileListRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	uid, code := util.ParseUidFromContext(appG.C)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	messageService := message_service.Message{
		ToUid:      uid,
		PageSize:   req.PageSize,
		PageNumber: req.PageNumber,
	}

	response := messageInboxProfileListResponse{MessageProfileList: []messageInboxProfile{}}
	messageList, messageUserList, err := messageService.GetMessageInboxProfileByToUid()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			appG.Response(http.StatusOK, code, response)
			return
		}
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
		return
	}

	var uids []int64
	for _, item := range *messageUserList {
		uids = append(uids, item.ToUid)
	}

	userService := user_service.User{}
	users, err := userService.GetUserprofilesByUids(uids)
	if err != nil {
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
		return
	}

	for idx := range *messageUserList {
		response.MessageProfileList = append(response.MessageProfileList, messageInboxProfile{
			ID:                (*messageUserList)[idx].ID,
			Mid:               (*messageUserList)[idx].Mid,
			Title:             (*messageList)[idx].Title,
			Time:              (*messageList)[idx].CreatedAt.Format("2006-01-02 15:04:05"),
			State:             (*messageUserList)[idx].State,
			FromUserName:      (*users)[idx].Name,
			FromUserAuthority: (*users)[idx].Authority,
			FromUserSeuID:     (*users)[idx].SeuID,
		})
	}

	totalNumber, err := messageService.GetTotalMessageCountByToUid()
	if err != nil {
		code = e.ERROR
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	response.TotalNumber = totalNumber

	appG.Response(http.StatusOK, code, response)
}

type messageInboxProfileDetailRequest struct {
	Mid int64 `json:"mid"`
}

type messageInboxProfileDetailResponse struct {
	Mid   int64  `json:"mid"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Time  string `json:"time"`

	FromUserName      string `json:"from_user_name"`
	FromUserAuthority int64  `json:"from_user_authority"`
	FromUserSeuID     string `json:"from_user_seu_id"`
}

func GetMessageInboxDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	var req messageInboxProfileDetailRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	uid, code := util.ParseUidFromContext(appG.C)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	messageService := message_service.Message{
		Mid:   req.Mid,
		ToUid: uid,
	}
	code = messageService.GetMessageInboxDetail()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	userService := user_service.User{
		UID: messageService.FromUid,
	}
	err := userService.GetUserProfileByUid()
	if err != nil {
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
		return
	}

	appG.Response(http.StatusOK, code, messageInboxProfileDetailResponse{
		Mid:               messageService.Mid,
		Title:             messageService.Title,
		Text:              messageService.Text,
		Time:              messageService.CreatedAt.Format("2006-01-02 15:04:05"),
		FromUserName:      userService.Name,
		FromUserAuthority: userService.Authority,
		FromUserSeuID:     userService.SeuID,
	})
}
