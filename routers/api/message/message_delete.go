package message

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/message_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type messageDeleteRequest struct {
	Mid int64 `json:"mid"`
}

func DeleteMessageAsUserView(c *gin.Context) {
	appG := app.Gin{C: c}
	var req messageDeleteRequest

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
	code = messageService.CheckIsReceiver()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	code = messageService.MarkAsUserDelete()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	appG.Response(http.StatusOK, code, nil)
}
