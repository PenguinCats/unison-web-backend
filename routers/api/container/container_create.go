package container

import (
	"github.com/PenguinCats/Unison-Elastic-Compute/api/types"
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/container_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ContainerCreateRequest struct {
	SlaveUUID       string  `json:"slave_uuid"`
	ImageName       string  `json:"image_name"`
	ExposedTcpPorts []int64 `json:"exposed_tcp_ports"`
	ExposedUdpPorts []int64 `json:"exposed_udp_ports"`
	CoreCnt         int     `json:"core_cnt"`
	MemSize         int64   `json:"mem_size"`
	StorageSize     int64   `json:"storage_size"`
}

type ContainerCreateResponse struct {
	Cid         int64 `json:"cid"`
	OperationID int64 `json:"operation_id"`
}

func ContainerCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	var req ContainerCreateRequest
	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	uid, code := util.ParseUidFromContext(appG.C)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	containerService := container_service.Container{
		Uid: uid,
	}
	operationID, code := containerService.Create(req.SlaveUUID, req.ImageName, req.ExposedTcpPorts, req.ExposedUdpPorts,
		req.CoreCnt, req.MemSize, req.StorageSize)

	resp := ContainerCreateResponse{
		Cid:         containerService.Cid,
		OperationID: operationID,
	}
	appG.Response(http.StatusOK, e.SUCCESS, resp)
}

func ContainerCreateCallback(c *gin.Context) {
	appG := app.Gin{C: c}
	var resp types.APIContainerCreateResponse
	if err := appG.C.BindJSON(&resp); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	cidInt64, err := strconv.ParseInt(resp.UECContainerID, 10, 64)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	containerService := container_service.Container{
		Cid: cidInt64,
	}
	_ = containerService.CreateCallback(resp.OperationID, resp.Code, resp.Msg)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
