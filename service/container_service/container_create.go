package container_service

import (
	"fmt"
	"github.com/PenguinCats/Unison-Elastic-Compute/api/types"
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/service/permission_group_service"
	"log"
	"strconv"
)

type Container struct {
	Cid int64 `json:"cid"`
	Hid int64 `json:"hid"`
	Uid int64 `json:"uid"`
	// extern info
	Ext string `json:"ext"`
}

func (c *Container) Create(uuid, imageName string, exposedTcpPorts, exposedUdpPorts []int64, coreCnt int,
	memSize, storageSize int64) (int64, int) {

	// step 1: check auth
	pus := permission_group_service.PermissionUserService{
		UID: c.Uid,
	}
	hosts, code := pus.GetHostsByUid()
	if code != e.SUCCESS {
		return -1, code
	}
	isAuth := false
	for _, item := range hosts {
		if item.UUID == uuid {
			isAuth = true
			c.Hid = item.Hid
			break
		}
	}
	if !isAuth {
		return -1, e.ERROR_AUTH_PERMISSION_DENIED
	}

	// step 2: create container record
	cid, err := models.AddContainer(models.Container{
		Hid: c.Hid,
		Uid: c.Uid,
		Ext: c.Ext,
	})
	if err != nil {
		return -1, e.ERROR
	}
	c.Cid = cid
	defer func() {
		if err != nil {
			_ = models.DeleteContainerByCID(cid)
		}
	}()

	// step 3: generate operation record
	operationID, err := models.AddOperation(c.Uid, fmt.Sprintf("UID [%d] create container [%d] on [%s]",
		c.Uid, cid, uuid))
	if err != nil {
		return -1, e.ERROR_OPERATION_GENERATION
	}

	// step 4: send create command
	var exposedTcpPortsString []string
	for p := range exposedTcpPorts {
		exposedTcpPortsString = append(exposedTcpPortsString, string(strconv.Itoa(p)))
	}
	var exposedUdpPortsString []string
	for p := range exposedUdpPorts {
		exposedUdpPortsString = append(exposedUdpPortsString, string(strconv.Itoa(p)))
	}
	callbackURL := setting.ServerSetting.CallbackPrefix + "/container/create_callback"
	req := types.APIContainerCreateRequest{
		APICallBackRequestBase: types.APICallBackRequestBase{
			OperationID: operationID,
			CallbackURL: callbackURL,
		},
		SlaveID:         uuid,
		UECContainerID:  strconv.FormatInt(c.Cid, 10),
		ImageName:       imageName,
		ExposedTCPPorts: exposedTcpPortsString,
		ExposedUDPPorts: exposedUdpPortsString,
		Mounts:          nil,
		CoreCnt:         coreCnt,
		MemorySize:      memSize,
		StorageSize:     storageSize,
	}

	OperationStatus := models.OperationStatusSentSuccess
	err = models.UpdateOperation(operationID, OperationStatus, "", "")
	if err != nil {
		log.Println(err.Error())
		return -1, e.ERROR
	}

	containerCreatePath := setting.UDCSetting.URL + "/container/create"
	errSendRequest := util.HttpPost(containerCreatePath, &req, nil)
	if errSendRequest != nil {
		// 必须和上面的 OperationStatusSentSuccess 分开，以防 callback 过早到来发生冲突
		OperationStatus = models.OperationStatusSentFail
		err = models.UpdateOperation(operationID, OperationStatus, "", "")
		if err != nil {
			log.Println(err.Error())
			return -1, e.ERROR
		}
	}

	return operationID, e.SUCCESS
}

func (c *Container) CreateCallback(operationID int64, code int, msg string) error {
	operationStatus, err := models.GetOperationStatus(operationID)
	if err != nil {
		return err
	}

	if operationStatus != models.OperationStatusSentSuccess {
		return nil
	}

	if code != types.SUCCESS {
		operationStatus = models.OperationStatusDoneFail
		_ = models.DeleteContainerByCID(c.Cid)
	} else {
		operationStatus = models.OperationStatusDoneSuccess
	}

	err = models.UpdateOperation(operationID, operationStatus, "", msg)
	if err != nil {
		log.Println(err.Error())
	}

	return err
}
