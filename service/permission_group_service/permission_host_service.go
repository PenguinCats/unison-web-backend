package permission_group_service

import (
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
)

type PermissionGroupHostService struct {
	ID      int64
	GroupID int64
	HIDs    []int64
}

func (pghs *PermissionGroupHostService) BindHosts() int {
	tx := models.NewContextForTransaction()

	err := models.DeletePermissionHostsFromGroup(tx, pghs.GroupID)
	if err != nil {
		tx.Rollback()
		return e.ERROR
	}

	err = models.AddPermissionHostsToGroup(tx, pghs.GroupID, pghs.HIDs)
	if err != nil {
		tx.Rollback()
		return e.ERROR
	}

	tx.Commit()
	return e.SUCCESS
}

func (pghs *PermissionGroupHostService) CountHosts() (map[int64]int64, int) {
	pghc, err := models.GetPermissionGroupHostsCount()
	if err != nil {
		return nil, e.ERROR
	}

	pghsMap := make(map[int64]int64)
	for _, item := range pghc {
		pghsMap[item.GroupID] = item.Cnt
	}
	return pghsMap, e.SUCCESS
}

func (pghs *PermissionGroupHostService) GetHosts() int {
	hosts, err := models.GetHostsIDByGroupID(pghs.GroupID)
	if err != nil {
		return e.ERROR
	}

	pghs.HIDs = hosts
	return e.SUCCESS
}
