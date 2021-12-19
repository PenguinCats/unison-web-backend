package models

import "time"

type Operation struct {
	OperationID int64 `gorm:"column:id; primaryKey"`
	Uid         int64 `gorm:"column:uid"`
	// 0: 指令待发出  1: 指令已发出  2: 发送指令失败  3: 指令执行成功  4: 指令执行出错
	Status      int64  `gorm:"column:status"`
	Description string `gorm:"column:description"`
	Result      string `gorm:"column:result"`

	CreatedAt time.Time `gorm:"column:time"`
}

type OperationStatus int64

const (
	OperationStatusNotSent OperationStatus = iota
	OperationStatusSentSuccess
	OperationStatusSentFail
	OperationStatusDoneSuccess
	OperationStatusDoneFail
	OperationException
)

func (os *OperationStatus) getInt64() int64 {
	return int64(*os)
}

func (Operation) TableName() string {
	return "operation"
}

func AddOperation(uid int64, description string) (int64, error) {
	operation := Operation{
		Uid:         uid,
		Status:      int64(OperationStatusNotSent),
		Description: description,
	}
	err := db.Create(&operation).Error
	return operation.OperationID, err
}

func UpdateOperation(operationId int64, status OperationStatus, description, result string) error {
	payload := map[string]interface{}{
		"status": status.getInt64(),
	}
	if description != "" {
		payload["description"] = description
	}
	if result != "" {
		payload["result"] = result
	}

	err := db.Model(&Operation{OperationID: operationId}).Updates(payload).Error
	return err
}

func GetOperationByUid(uid int64) (ops []Operation, err error) {
	err = db.Where("uid = ?", uid).Find(&ops).Error
	return
}

func GetOperationStatus(operationID int64) (OperationStatus, error) {
	var operation Operation
	err := db.Find(&operation, operationID).Error
	if err != nil {
		return OperationException, err
	}
	return OperationStatus(operation.Status), nil
}
