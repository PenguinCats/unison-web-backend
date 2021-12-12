package host

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/host_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateHostExtRequest struct {
	Hid int64  `json:"hid"`
	Ext string `json:"ext"`
}

func UpdateHostExt(c *gin.Context) {
	appG := app.Gin{C: c}
	var req updateHostExtRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	hs := host_service.HostService{
		Hid: req.Hid,
		Ext: req.Ext,
	}
	code := hs.UpdateExt()
	appG.Response(http.StatusOK, code, nil)
}
