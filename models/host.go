package models

type Host struct {
	Hid  int64  `gorm:"column:hid; primaryKey"`
	UUID string `gorm:"column:uuid"`
	// extern info
	Ext string `gorm:"column:ext"`
}

func (Host) TableName() string {
	return "host"
}

func GetHostAll() (*[]Host, error) {
	var hosts []Host
	err := db.Find(&hosts).Error
	if err != nil {
		return nil, err
	}

	return &hosts, nil
}

func AddHost(h Host) error {
	err := db.Create(&h).Error
	return err
}

func UpdateExt(hid int64, ext string) error {
	h := Host{
		Hid: hid,
		Ext: ext,
	}
	err := db.Model(&h).Update("ext", ext).Error
	return err
}

func DeleteHostByUUID(uuid string) error {
	err := db.Where("uuid = ?", uuid).Delete(&Host{}).Error
	return err
}
