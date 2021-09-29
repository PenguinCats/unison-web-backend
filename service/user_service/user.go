package user_service

import (
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"log"
)

type User struct {
	UID       int64
	Username  string
	Password  string
	Name      string
	Authority int64
	SeuID     string
}

func (u *User) FillUidByUserName() error {
	uid, err := models.GetUIDByUserName(u.Username)
	if err != nil {
		return err
	}

	u.UID = uid
	return nil
}

func (u *User) FillAuthorityByUid() error {
	authority, err := models.GetAuthByUID(u.UID)
	if err != nil {
		return err
	}

	u.Authority = authority
	return nil
}

func (u *User) CheckPassword(pwdSalt, salt string) bool {
	if u.UID == 0 {
		log.Println("uid is not fetched")
		return false
	}
	pwd, err := models.GetPwdByUid(u.UID)
	if err != nil {
		return false
	}

	pwdWithSalt := util.EncodeSHA512(pwd + salt)

	if pwdSalt != pwdWithSalt {
		return false
	}

	return true
}

func (u *User) GetUserByUid() error {
	user, err := models.GetUserByUID(u.UID)
	if err != nil {
		return err
	}

	u.Username = user.Username
	u.Name = user.Name
	u.Authority = user.Authority
	u.SeuID = user.SeuID

	return nil
}
