package auth_service

import (
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
)

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

func (a *Auth) IsSelfOrRoot(outerUID int64) bool {
	if a.Uid == outerUID || a.IsRoot() {
		return true
	}
	return false
}

func (a *Auth) IsSelfOrAdmin(outerUID int64) bool {
	if a.Uid == outerUID || a.IsAdmin() {
		return true
	}
	return false
}

func (a *Auth) GetAuthority() (int64, int) {
	auth, err := models.GetAuthByUID(a.Uid)
	if err != nil {
		return auth, e.ERROR
	}
	return auth, e.SUCCESS
}
