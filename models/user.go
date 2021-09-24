package models

import "gorm.io/gorm"

type User struct {
	UID       int64  `gorm:"column:uid, primaryKey"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	Name      string `gorm:"column:name"`
	Authority int64  `gorm:"column:authority"`
	SeuID     string `gorm:"column:seu_id"`
}

func (User) TableName() string {
	return "user"
}

func GetUIDByUserName(username string) (int64, error) {
	//var user User
	//user.Username = username
	var uid int64
	err := db.Model(&User{}).Select("uid").Where("username = ?", username).Take(&uid).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	//return user.UID, nil
	return uid, nil
}

func ExistUserByUserName(username string) (bool, error) {
	uid, err := GetUIDByUserName(username)
	if err != nil || uid <= 0 {
		return false, err
	}
	return true, nil
}

func GetUserByUID(uid int64) (*User, error) {
	var user User
	user.UID = uid
	err := db.Where("uid = ?", uid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserAll() ([]*User, error) {
	var users []*User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetAuthByUID(uid int64) (int64, error) {
	var auth int64
	err := db.Model(&User{}).Select("authority").Where("uid = ?", uid).Take(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return auth, nil
}
