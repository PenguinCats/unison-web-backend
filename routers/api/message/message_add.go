package message

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/message_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type messageAddRequest struct {
	To    []int64 `json:"to"`
	Title string  `json:"title"`
	Text  string  `json:"text"`
	From  int64   `json:"from"`
}

func AddMessage(c *gin.Context) {
	appG := app.Gin{C: c}
	var req messageAddRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	messageService := message_service.Message{
		Title:   req.Title,
		Text:    req.Text,
		FromUid: req.From,
		ToUids:  req.To,
	}
	err := messageService.AddMessage()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
