package host

import (
	"github.com/PenguinCats/Unison-Elastic-Compute/api/types"
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/host_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HostDeleteRequest struct {
	HostUUID string `json:"host_uuid"`
}

func DeleteHost(c *gin.Context) {
	appG := app.Gin{C: c}
	var req HostDeleteRequest

	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	uid, code := util.ParseUidFromContext(appG.C)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	hostService := host_service.HostService{
		UUID: req.HostUUID,
	}

	code = hostService.DeleteHost(uid)
	appG.Response(http.StatusOK, code, nil)
}

func DeleteHostCallback(c *gin.Context) {
	appG := app.Gin{C: c}
	var resp types.APISlaveDeleteResponse
	if err := appG.C.BindJSON(&resp); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	hostService := host_service.HostService{
		UUID: resp.SlaveUUID,
	}
	_ = hostService.DeleteHostCallback(resp.OperationID, resp.Code, resp.Msg)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
