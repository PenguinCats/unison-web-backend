package auth_service

import "github.com/PenguinCats/unison-web-backend/models"

type Auth struct {
	Uid int64

	// log in by username and password
	UserName string
	Password string

	// log in by SEU ID
	SEUID string

	// log in by QQ ID
	QQID string
}

func (a *Auth) IsRoot() bool {
	auth, err := models.GetAuthByUID(a.Uid)
	if err != nil || auth != 1 {
		return false
	}
	return true
}

func (a *Auth) IsAdmin() bool {
	auth, err := models.GetAuthByUID(a.Uid)
	if err != nil || (auth != 1 && auth != 2) {
		return false
	}
	return true
}
