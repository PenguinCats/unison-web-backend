package operation

import (
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/service/operation_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OperationStatusRequest struct {
	ID int64 `json:"id"`
}

type OperationStatusResponse struct {
	Status int64 `json:"status"`
}

func OperationStatus(c *gin.Context) {
	appG := app.Gin{C: c}
	var req OperationStatusRequest
	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	ops := operation_service.Operation{
		OperationID: req.ID,
	}

	status, code := ops.GetStatus()
	if code != e.SUCCESS {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	resp := OperationStatusResponse{
		Status: status,
	}
	appG.Response(http.StatusOK, e.SUCCESS, resp)
}
