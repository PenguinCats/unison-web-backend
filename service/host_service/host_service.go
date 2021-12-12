package host_service

import (
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
)

type HostService struct {
	Hid  int64  `json:"hid"`
	UUID string `json:"uuid"`
	// extern info
	Ext string `json:"ext"`
}

func (hs *HostService) UpdateExt() int {
	err := models.UpdateExt(hs.Hid, hs.Ext)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}
