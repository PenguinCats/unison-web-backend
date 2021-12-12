package host_service

import (
	"fmt"
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"testing"
)

func TestHostDaemon_CheckUDCHostList(t *testing.T) {
	setting.UDCSetting.URL = "http://127.0.0.1:9600/api"
	models.SetupForTest()
	SetUp()

	err := hostD.checkUDCHostList()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestHostDaemon_GetAllHostList(t *testing.T) {
	setting.UDCSetting.URL = "http://127.0.0.1:9600/api"
	models.SetupForTest()
	SetUp()

	profile, host, lost, err := hostD.GetAllHostList()
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(profile)
	fmt.Println(host)
	fmt.Println(lost)
}
