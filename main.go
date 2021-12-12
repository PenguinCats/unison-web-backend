package main

import (
	"fmt"
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"github.com/PenguinCats/unison-web-backend/routers"
	"github.com/PenguinCats/unison-web-backend/service/host_service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
	util.Setup()

	host_service.SetUp()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
