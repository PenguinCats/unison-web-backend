package host

import (
	"github.com/PenguinCats/Unison-Elastic-Compute/api/types"
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HostAuthorizedListItem struct {
	HID      int64  `json:"hid"`
	HostUUId string `json:"host_uuid"`
	Ext      string `json:"ext"`

	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platform_family"`
	PlatformVersion string `json:"platform_version"`
	MemoryTotalSize uint64 `json:"memory_total_size"`
	CpuModelName    string `json:"cpu_model_name"`
	LogicalCoreCnt  int    `json:"logical_core_cnt"`
	PhysicalCoreCnt int    `json:"physical_core_cnt"`
}

type HostAuthorizedListResponse struct {
	Hosts []HostAuthorizedListItem `json:"hosts"`
}

func GetHostAuthorizedList(c *gin.Context) {
	appG := app.Gin{C: c}

	uid, code := util.ParseUidFromContext(appG.C)
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	pus := permission_group_service.PermissionUserService{
		UID: uid,
	}

	hosts, code := pus.GetHostsByUid()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}
	var uuids []string
	for _, host := range hosts {
		uuids = append(uuids, host.UUID)
	}

	hostProfileListPath := setting.UDCSetting.URL + "/slave/profile"
	req := types.APISlaveProfileListRequest{SlavesUUID: uuids}
	var res types.APISlaveProfileListResponse
	err := util.HttpPost(hostProfileListPath, &req, &res)
	if err != nil {
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
	}

	if len(res.Slaves) != len(hosts) {
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
	}

	response := HostAuthorizedListResponse{Hosts: []HostAuthorizedListItem{}}

	for idx, host := range hosts {
		// 若远程返回的结果 UUID 字段为空，则表示无法识别，已丢失
		if res.Slaves[idx].SlaveUUId == host.UUID {
			response.Hosts = append(response.Hosts, HostAuthorizedListItem{
				HID:             host.Hid,
				HostUUId:        host.UUID,
				Ext:             host.Ext,
				Platform:        res.Slaves[idx].Platform,
				PlatformFamily:  res.Slaves[idx].PlatformFamily,
				PlatformVersion: res.Slaves[idx].PlatformVersion,
				MemoryTotalSize: res.Slaves[idx].MemoryTotalSize,
				CpuModelName:    res.Slaves[idx].CpuModelName,
				LogicalCoreCnt:  res.Slaves[idx].LogicalCoreCnt,
				PhysicalCoreCnt: res.Slaves[idx].PhysicalCoreCnt,
			})
		}
	}

	appG.Response(http.StatusOK, code, response)
}
