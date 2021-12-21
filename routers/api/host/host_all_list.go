package host

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/host_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HostAllListItem struct {
	HID      int64  `json:"hid"`
	HostUUId string `json:"host_uuid"`
	Ext      string `json:"ext"`

	IsLost bool `json:"is_lost"`

	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platform_family"`
	PlatformVersion string `json:"platform_version"`
	MemoryTotalSize uint64 `json:"memory_total_size"`
	CpuModelName    string `json:"cpu_model_name"`
	LogicalCoreCnt  int    `json:"logical_core_cnt"`
	PhysicalCoreCnt int    `json:"physical_core_cnt"`
}

type HostAllListResponse struct {
	Hosts []HostAllListItem `json:"hosts"`
}

func GetHostAllList(c *gin.Context) {
	appG := app.Gin{C: c}

	profile, host, lost, err := host_service.GetHostDaemon().GetAllHostList()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
	}

	var hosts []HostAllListItem
	for idx := range host {
		hosts = append(hosts, HostAllListItem{
			HID:             host[idx].Hid,
			HostUUId:        host[idx].UUID,
			Ext:             host[idx].Ext,
			IsLost:          false,
			Platform:        profile[idx].Platform,
			PlatformFamily:  profile[idx].PlatformFamily,
			PlatformVersion: profile[idx].PlatformVersion,
			MemoryTotalSize: profile[idx].MemoryTotalSize,
			CpuModelName:    profile[idx].CpuModelName,
			LogicalCoreCnt:  profile[idx].LogicalCoreCnt,
			PhysicalCoreCnt: profile[idx].PhysicalCoreCnt,
		})
	}
	for _, host := range lost {
		hosts = append(hosts, HostAllListItem{
			HID:      host.Hid,
			HostUUId: host.UUID,
			Ext:      host.Ext,
			IsLost:   true,
		})
	}

	appG.Response(http.StatusOK, e.SUCCESS, &HostAllListResponse{Hosts: hosts})
}
