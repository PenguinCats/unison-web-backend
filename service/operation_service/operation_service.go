package operation_service

import (
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"time"
)

type Operation struct {
	OperationID int64 `gorm:"column:id; primaryKey"`
	Uid         int64 `gorm:"column:uid"`
	// 0: 指令待发出  1: 指令已发出  2: 发送指令失败  3: 指令执行成功  4: 指令执行出错
	Status      int64  `gorm:"column:status"`
	Description string `gorm:"column:description"`
	Result      string `gorm:"column:result"`

	CreatedAt time.Time `gorm:"column:time"`
}

func (ope *Operation) GetStatus() (int64, int) {
	status, err := models.GetOperationStatus(ope.OperationID)
	if err != nil {
		return 0, e.ERROR
	}

	return status.GetInt64(), e.SUCCESS
}
