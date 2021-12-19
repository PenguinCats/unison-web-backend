package host_service

import (
	"fmt"
	"github.com/PenguinCats/Unison-Elastic-Compute/api/types"
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"log"
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

func (hs *HostService) DeleteHost(uid int64) int {
	operationID, err := models.AddOperation(uid, fmt.Sprintf("UID [%d] delete host [%s]", uid, hs.UUID))
	if err != nil {
		return e.ERROR_OPERATION_GENERATION
	}

	callbackURL := setting.ServerSetting.CallbackPrefix + "/host/delete_callback"
	req := types.APISlaveDeleteRequest{
		APICallBackRequestBase: types.APICallBackRequestBase{
			OperationID: operationID,
			CallbackURL: callbackURL,
		},
		SlaveUUID: hs.UUID,
	}

	OperationStatus := models.OperationStatusSentSuccess
	errUnexpected := models.UpdateOperation(operationID, OperationStatus, "", "")
	if errUnexpected != nil {
		log.Println(errUnexpected.Error())
	}

	hostDeletePath := setting.UDCSetting.URL + "/slave/delete"
	errSendRequest := util.HttpPost(hostDeletePath, &req, nil)
	if errSendRequest != nil {
		// 必须和上面的 OperationStatusSentSuccess 分开，以防 callback 过早到来发生冲突
		OperationStatus = models.OperationStatusSentFail
		errUnexpected = models.UpdateOperation(operationID, OperationStatus, "", "")
		if errUnexpected != nil {
			log.Println(errUnexpected.Error())
		}
	}

	return e.SUCCESS
}

func (hs *HostService) DeleteHostCallback(operationID int64, code int, msg string) error {
	operationStatus, err := models.GetOperationStatus(operationID)
	if err != nil {
		return err
	}

	if operationStatus != models.OperationStatusSentSuccess {
		return nil
	}

	if code != types.SUCCESS {
		operationStatus = models.OperationStatusDoneFail
	} else {
		operationStatus = models.OperationStatusDoneSuccess
	}

	err = models.UpdateOperation(operationID, operationStatus, "", msg)
	if err != nil {
		log.Println(err.Error())
	}

	GetHostDaemon().RemoveHostUUIDFromDaemon(hs.UUID)
	err = models.DeleteHostByUUID(hs.UUID)

	return err
}
