package models

import "gorm.io/gorm"

type PermissionGroup struct {
	GroupID int64  `gorm:"column:group_id; primaryKey"`
	Name    string `gorm:"column:name"`
}

func (PermissionGroup) TableName() string {
	return "permission_group"
}

func GetPermissionNameByID(id int64) (string, error) {
	var name string
	err := db.Model(&PermissionGroup{}).Select("name").Where("group_id = ?", id).Take(&name).Error
	return name, err
}

func GetPermissionGroupsByIDs(ids []int64) (*[]PermissionGroup, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var groups []PermissionGroup

	err := db.Model(&PermissionGroup{}).Where("group_id IN ?", ids).Find(&groups).Error
	return &groups, err
}

func GetPermissionGroupList() (*[]PermissionGroup, error) {
	var groups []PermissionGroup
	err := db.Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return &groups, nil
}

func UpdatePermissionGroupName(id int64, name string) error {
	pg := PermissionGroup{
		GroupID: id,
		Name:    name,
	}
	err := db.Save(&pg).Error
	//err := db.Model(&PermissionGroup{}).Update(&PermissionGroup{
	//	GroupID: id,
	//	Name:    name,
	//})
	return err
}

func DeletePermissionGroup(id int64) error {
	err := db.Delete(&PermissionGroup{}, id).Error
	return err
}

func AddPermissionGroup(name string) error {
	err := db.Create(&PermissionGroup{
		Name: name,
	}).Error
	return err
}

type PermissionUser struct {
	ID      int64 `gorm:"column:id; primaryKey"`
	GroupID int64 `gorm:"column:group_id"`
	UID     int64 `gorm:"column:uid"`
}

func (PermissionUser) TableName() string {
	return "permission_user"
}

func GetPermissionGroupIDByUid(uid int64) (*[]int64, error) {
	var id []int64
	err := db.Model(&PermissionUser{}).Select("group_id").Where("uid = ?", uid).Find(&id).Error
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func AddPermissionUser(tx *gorm.DB, uid int64, groupIdList []int64) error {
	if len(groupIdList) == 0 {
		return nil
	}
	var puList []PermissionUser
	for _, gid := range groupIdList {
		puList = append(puList, PermissionUser{
			GroupID: gid,
			UID:     uid,
		})
	}
	err := tx.Create(&puList).Error
	return err
}

func DeletePermissionUser(tx *gorm.DB, uid int64) error {
	err := tx.Where("uid = ?", uid).Delete(&PermissionUser{}).Error
	return err
}

type PermissionHost struct {
	ID      int64 `gorm:"column:id; primaryKey"`
	GroupID int64 `gorm:"column:group_id"`
	HID     int64 `gorm:"column:hid"`
}

func GetHostsIDByGroupID(gid int64) (id []int64, err error) {
	err = db.Model(&PermissionHost{}).Select("hid").Where("group_id = ?", gid).Find(&id).Error
	if err != nil {
		return nil, err
	}
	return
}

func (PermissionHost) TableName() string {
	return "host_permission"
}

func AddPermissionHostsToGroup(tx *gorm.DB, gid int64, hosts []int64) error {
	if len(hosts) == 0 {
		return nil
	}
	var phList []PermissionHost
	for _, hid := range hosts {
		phList = append(phList, PermissionHost{
			GroupID: gid,
			HID:     hid,
		})
	}
	err := tx.Create(&phList).Error
	return err
}

func DeletePermissionHostsFromGroup(tx *gorm.DB, gid int64) error {
	err := tx.Where("group_id = ?", gid).Delete(&PermissionHost{}).Error
	return err
}

type PermissionGroupHostsCount struct {
	GroupID int64 `gorm:"column:group_id"`
	Cnt     int64 `gorm:"column:cnt"`
}

func GetPermissionGroupHostsCount() ([]PermissionGroupHostsCount, error) {
	var pghc []PermissionGroupHostsCount
	err := db.Table(PermissionHost{}.TableName()).Select("group_id", "count(*) as cnt").Group("group_id").Scan(&pghc).Error
	if err != nil {
		return nil, err
	}
	return pghc, nil
}
