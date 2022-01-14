package host

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/host_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HostProfileRequest struct {
	UUID string `json:"uuid"`
}

func GetHostProfile(c *gin.Context) {
	appG := app.Gin{C: c}
	var request HostProfileRequest
	if err := appG.C.BindJSON(&request); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	slaves, err := host_service.GetHostDaemon().GetHostProfile([]string{request.UUID})
	if err != nil || len(slaves) != 1 || slaves[0].SlaveUUId != request.UUID {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	response := HostprofileItem{
		HostUUId:        slaves[0].SlaveUUId,
		Platform:        slaves[0].Platform,
		PlatformFamily:  slaves[0].PlatformFamily,
		PlatformVersion: slaves[0].PlatformVersion,
		MemoryTotalSize: slaves[0].MemoryTotalSize,
		CpuModelName:    slaves[0].CpuModelName,
		LogicalCoreCnt:  slaves[0].LogicalCoreCnt,
		PhysicalCoreCnt: slaves[0].PhysicalCoreCnt,
	}
	appG.Response(http.StatusOK, e.SUCCESS, &response)
}
