package host

import (
	"github.com/PenguinCats/Unison-Elastic-Compute/api/types"
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HostStatusRequest struct {
	UUIDList []string `json:"uuids"`
}

type HostStatusItem struct {
	UUID             string `json:"uuid"`
	Stats            string `json:"stats"`
	CoreAvailable    int    `json:"core_available"`
	MemAvailable     uint64 `json:"mem_available"`
	StorageAvailable uint64 `json:"storage_available"`
}

type HostStatusResponse struct {
	Status []HostStatusItem `json:"status"`
}

func GetHostStatus(c *gin.Context) {
	appG := app.Gin{C: c}
	var req HostStatusRequest
	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	code := e.SUCCESS

	hostProfileListPath := setting.UDCSetting.URL + "/slave/status"
	RemoteReq := types.APISlaveStatusRequest{SlaveUUID: req.UUIDList}
	var res types.APISlaveStatusResponse
	err := util.HttpPost(hostProfileListPath, &RemoteReq, &res)
	if err != nil {
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
		return
	}

	response := HostStatusResponse{Status: []HostStatusItem{}}
	for _, status := range res.Status {
		response.Status = append(response.Status, HostStatusItem{
			UUID:             status.UUID,
			Stats:            status.Stats,
			CoreAvailable:    status.CoreAvailable,
			MemAvailable:     status.MemAvailable,
			StorageAvailable: status.StorageAvailable,
		})
	}

	appG.Response(http.StatusOK, code, response)
}
