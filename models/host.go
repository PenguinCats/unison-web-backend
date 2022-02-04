package models

import "gorm.io/gorm"

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

func UpdateHostExt(hid int64, ext string) error {
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

func GetHostUUIDByHid(hid []int64) (uuids []Host, err error) {
	if len(hid) == 0 {
		return []Host{}, nil
	}
	err = db.Where("hid in ?", hid).Find(&uuids).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []Host{}, nil
		}
		return nil, err
	}
	return
}
