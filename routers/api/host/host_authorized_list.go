package host

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/host_service"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HostprofileItem struct {
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
	Hosts []HostprofileItem `json:"hosts"`
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

	slaves, err := host_service.GetHostDaemon().GetHostProfile(uuids)
	if err != nil {
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
		return
	}

	if len(slaves) != len(hosts) {
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
		return
	}

	response := HostAuthorizedListResponse{Hosts: []HostprofileItem{}}

	for idx, host := range hosts {
		// 若远程返回的结果 UUID 字段为空，则表示无法识别，已丢失
		if slaves[idx].SlaveUUId == host.UUID {
			response.Hosts = append(response.Hosts, HostprofileItem{
				HID:             host.Hid,
				HostUUId:        host.UUID,
				Ext:             host.Ext,
				Platform:        slaves[idx].Platform,
				PlatformFamily:  slaves[idx].PlatformFamily,
				PlatformVersion: slaves[idx].PlatformVersion,
				MemoryTotalSize: slaves[idx].MemoryTotalSize,
				CpuModelName:    slaves[idx].CpuModelName,
				LogicalCoreCnt:  slaves[idx].LogicalCoreCnt,
				PhysicalCoreCnt: slaves[idx].PhysicalCoreCnt,
			})
		}
	}

	appG.Response(http.StatusOK, code, response)
}
