package permission_group_service

import (
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"gorm.io/gorm"
)

type PermissionGroupService struct {
	GroupID int64
	Name    string
}

func (p *PermissionGroupService) GetGroupNameByGroupID() (string, int) {
	name, err := models.GetPermissionNameByID(p.GroupID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return name, e.SUCCESS
		}
		return "", e.ERROR
	}

	return name, e.SUCCESS
}
