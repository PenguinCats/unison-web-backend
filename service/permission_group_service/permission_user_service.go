package permission_group_service

import (
	"fmt"
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"log"
)

type PermissionUserService struct {
	ID      int64
	GroupID int64
	UID     int64
}

func (p *PermissionUserService) GetGroupIDsByUid() (*[]int64, int) {
	ids, err := models.GetPermissionGroupIDByUid(p.UID)
	if err != nil {
		return nil, e.ERROR
	}

	return &ids, e.SUCCESS
}

func (p *PermissionUserService) GetHostsByUid() ([]models.Host, int) {
	gids, err := models.GetPermissionGroupIDByUid(p.UID)
	if err != nil {
		return nil, e.ERROR
	}

	hids, err := models.GetHostsIDByGroupIDs(gids)
	if err != nil {
		return nil, e.ERROR
	}

	hosts, err := models.GetHostUUIDByHid(hids)
	// 为保证有序
	mp := map[int64]models.Host{}
	for _, host := range hosts {
		mp[host.Hid] = host
	}

	var hostsRet []models.Host
	for _, hid := range hids {
		if uuid, ok := mp[hid]; ok {
			hostsRet = append(hostsRet, uuid)
		} else {
			log.Println(fmt.Sprintf("cannot find uuid for hid [%d]", hid))
		}
	}
	return hostsRet, e.SUCCESS
}
