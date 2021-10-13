package user_service

import (
	"errors"
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

	Query string

	PageSize   int64
	PageNumber int64
}

func (u *User) GetUserList() (*[]models.User, error) {
	offset := u.PageSize * (u.PageNumber - 1)
	if offset < 0 {
		offset = 0
	}
	if u.PageSize <= 0 {
		u.PageSize = 10
	}

	userList, err := models.GetUserList(int(offset), int(u.PageSize))

	return userList, err
}

func (u *User) GetTotalUser() (int64, error) {
	return models.GetUserCount()
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

func (u *User) GetUserProfileByUid() error {
	user, err := models.GetUserProfileByUID(u.UID)
	if err != nil {
		return err
	}

	u.Name = user.Name
	u.Authority = user.Authority
	u.SeuID = user.SeuID

	return nil
}

func (u *User) GetUserProfilesByUids(uids []int64) (*[]models.User, error) {
	users, err := models.GetUserProfilesByUIDs(uids)
	if err != nil {
		return nil, err
	}

	mp := make(map[int64]models.User)
	for _, item := range *users {
		mp[item.UID] = item
	}

	var usersRet []models.User
	for _, uid := range uids {
		val, ok := mp[uid]
		if !ok {
			return nil, errors.New("user broken")
		}
		usersRet = append(usersRet, val)
	}

	return &usersRet, nil
}

func (u *User) GetUserProfileByQueryString() (*[]models.User, error) {
	return models.GetUserProfilesByQuery(u.Query)
}
