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
			return "", e.SUCCESS
		}
		return "", e.ERROR
	}

	return name, e.SUCCESS
}

func (p *PermissionGroupService) GetPermissionGroupList() (*[]models.PermissionGroup, int) {
	users, err := models.GetPermissionGroupList()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, e.SUCCESS
		}
		return nil, e.ERROR
	}
	return users, e.SUCCESS
}

func (p *PermissionGroupService) GetGroupsByGroupIDs(ids []int64) (*[]models.PermissionGroup, int) {
	groups, err := models.GetPermissionGroupsByIDs(ids)
	if err != nil {
		return nil, e.ERROR

	}
	return groups, e.SUCCESS
}

func (p *PermissionGroupService) EditGroupName() int {
	err := models.UpdatePermissionGroupName(p.GroupID, p.Name)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

func (p *PermissionGroupService) DeleteGroup() int {
	err := models.DeletePermissionGroup(p.GroupID)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

func (p *PermissionGroupService) AddGroup() int {
	err := models.AddPermissionGroup(p.Name)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}
