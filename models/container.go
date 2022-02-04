package models

type Container struct {
	Cid int64 `gorm:"column:cid; primaryKey"`
	Hid int64 `gorm:"column:hid"`
	Uid int64 `gorm:"column:uid"`
	// extern info
	Ext string `gorm:"column:ext"`
}

func (Container) TableName() string {
	return "container"
}

func GetContainerAll() (*[]Container, error) {
	var containers []Container
	err := db.Find(&containers).Error
	if err != nil {
		return nil, err
	}

	return &containers, nil
}

func GetContainerByHid(hid int64) (*[]Container, error) {
	var containers []Container
	err := db.Where("hid = ?", hid).Find(&containers).Error
	if err != nil {
		return nil, err
	}

	return &containers, nil
}

func GetContainerByUid(uid int64) (*[]Container, error) {
	var containers []Container
	err := db.Where("uid = ?", uid).Find(&containers).Error
	if err != nil {
		return nil, err
	}

	return &containers, nil
}

func AddContainer(c Container) (int64, error) {
	err := db.Create(&c).Error
	return c.Cid, err
}

func UpdateContainerExt(cid int64, ext string) error {
	h := Container{
		Cid: cid,
		Ext: ext,
	}
	err := db.Model(&h).Update("ext", ext).Error
	return err
}

func DeleteContainerByCID(cid int64) error {
	err := db.Where("cid = ?", cid).Delete(&Container{}).Error
	return err
}
