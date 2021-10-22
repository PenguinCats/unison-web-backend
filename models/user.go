package models

import "gorm.io/gorm"

type User struct {
	UID       int64  `gorm:"column:uid; primaryKey"`
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
	if err != nil {
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

func GetUserCount() (int64, error) {
	var count int64
	err := db.Model(&User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetUserList(offset, limit int) (*[]User, error) {
	var users []User
	err := db.Select("uid", "username", "name", "authority", "seu_id").Offset(offset).
		Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func GetAuthByUID(uid int64) (int64, error) {
	var auth int64
	err := db.Model(&User{}).Select("authority").Where("uid = ?", uid).Take(&auth).Error
	if err != nil {
		return 0, err
	}

	return auth, nil
}

func GetPwdByUid(uid int64) (string, error) {
	var pwd string
	err := db.Model(&User{}).Select("password").Where("uid = ?", uid).Take(&pwd).Error
	if err != nil {
		return "", err
	}

	return pwd, nil
}

func GetUserProfileByUID(uid int64) (*User, error) {
	var user User
	err := db.Model(&User{}).Select("uid", "name", "username", "authority", "seu_id").
		Where("uid = ?", uid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserProfilesByUIDs(uids []int64) (*[]User, error) {
	var users []User
	err := db.Model(&User{}).Select("uid", "name", "authority", "seu_id").
		Where("uid IN ?", uids).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func GetUserProfilesByQuery(query string) (*[]User, error) {
	query = "%" + query + "%"

	var users []User
	err := db.Model(&User{}).Select("uid", "name", "authority", "seu_id").
		Where("name LIKE ?", query).Or("seu_id LIKE ?", query).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func DeleteUser(uid int64) error {
	return db.Delete(&User{}, uid).Error
}

func UserAdd(tx *gorm.DB, u User) (*User, error) {
	err := tx.Create(&u).Error
	return &u, err
}

func UserEdit(tx *gorm.DB, u User) error {
	err := tx.Model(&u).Updates(User{
		Name:      u.Name,
		Authority: u.Authority,
		SeuID:     u.SeuID,
	}).Error
	return err
}

func EditUserPassword(uid int64, pwd string) error {
	err := db.Model(&User{UID: uid}).Update("password", pwd).Error
	return err
}
