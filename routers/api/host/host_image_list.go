package host

import (
	"github.com/PenguinCats/Unison-Elastic-Compute/api/types"
	"github.com/PenguinCats/unison-web-backend/pkg/app"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HostImageListRequest struct {
	UUID string `json:"uuid"`
}

type ImageListItem struct {
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	CreatedTime string `json:"created_time"`
}

type HostImageListResponse struct {
	Images []ImageListItem `json:"images"`
}

func GetHostImageList(c *gin.Context) {
	appG := app.Gin{C: c}
	var req HostImageListRequest
	if err := appG.C.BindJSON(&req); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	code := e.SUCCESS

	hostProfileListPath := setting.UDCSetting.URL + "/slave/image_list"
	RemoteReq := types.APISlaveImageListRequest{SlaveUUID: req.UUID}
	var res types.APISlaveImageListResponse
	err := util.HttpPost(hostProfileListPath, &RemoteReq, &res)
	if err != nil {
		code = e.ERROR
		appG.Response(http.StatusOK, code, nil)
		return
	}

	response := HostImageListResponse{Images: []ImageListItem{}}
	for _, img := range res.Images {
		response.Images = append(response.Images, ImageListItem{
			Name:        img.Name,
			Size:        img.Size,
			CreatedTime: img.CreatedTime,
		})
	}

	appG.Response(http.StatusOK, code, response)
}
