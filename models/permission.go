package models

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
