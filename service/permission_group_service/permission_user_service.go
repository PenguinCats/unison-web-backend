package permission_group_service

import (
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"gorm.io/gorm"
)

type PermissionUserService struct {
	ID      int64
	GroupID int64
	UID     int64
}

func (p *PermissionUserService) GetGroupIDsByUid() (*[]int64, int) {
	ids, err := models.GetPermissionGroupIDByUid(p.UID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &[]int64{}, e.SUCCESS
		}
		return nil, e.ERROR
	}

	return ids, e.SUCCESS
}
