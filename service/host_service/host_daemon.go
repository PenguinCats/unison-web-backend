package host_service

import (
	"github.com/PenguinCats/Unison-Elastic-Compute/api/types"
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"github.com/PenguinCats/unison-web-backend/pkg/util"
	"golang.org/x/sync/singleflight"
	"log"
	"time"
)

type HostDaemon struct {
	ExistingHostUUIDList map[string]struct{}

	*singleflight.Group
}

var hostD *HostDaemon

func SetUp() {
	hostD = &HostDaemon{
		ExistingHostUUIDList: make(map[string]struct{}),
		Group:                &singleflight.Group{},
	}

	err := hostD.loadExistingHost()
	if err != nil {
		log.Fatalf(err.Error())
	}

	hostD.startHostListCheckPeriodicity()
}

func GetHostDaemon() *HostDaemon {
	return hostD
}

func (hd *HostDaemon) loadExistingHost() error {
	hosts, err := models.GetHostAll()
	if err != nil {
		return err
	}
	for _, host := range *hosts {
		hd.ExistingHostUUIDList[host.UUID] = struct{}{}
	}
	return nil
}

func (hd *HostDaemon) startHostListCheckPeriodicity() {
	go func() {
		t := time.NewTicker(time.Minute * 1)
		for {
			select {
			case <-t.C:
				_ = hd.checkUDCHostList()
			}
		}
	}()
}

func (hd *HostDaemon) checkUDCHostList() error {
	_, err, _ := hd.Do("checkHostList", func() (interface{}, error) {
		hostListPath := setting.UDCSetting.URL + "/slave/uuid_list"

		var res types.APISlaveUUIDListResponse
		err := util.HttpPost(hostListPath, nil, &res)
		if err != nil {
			return nil, err
		}

		for _, slave := range res.SlavesUUID {
			if _, ok := hd.ExistingHostUUIDList[slave]; !ok {
				err := models.AddHost(models.Host{
					UUID: slave,
				})
				if err != nil {
					log.Println(err.Error())
				} else {
					hd.ExistingHostUUIDList[slave] = struct{}{}
				}
			}
		}
		return nil, nil
	})

	return err
}

func (hd *HostDaemon) GetAllHostList() (hostProfile []types.SlaveProfile, host []models.Host,
	lost []models.Host, err error) {

	err = hd.checkUDCHostList()
	if err != nil {
		return nil, nil, nil, err
	}

	allHosts, err := models.GetHostAll()
	if err != nil {
		return nil, nil, nil, err
	}

	var allHostUUIDList []string
	uuid2host := map[string]models.Host{}
	for _, host := range *allHosts {
		allHostUUIDList = append(allHostUUIDList, host.UUID)
		uuid2host[host.UUID] = host
	}

	hostProfileListPath := setting.UDCSetting.URL + "/slave/profile"
	req := types.APISlaveProfileListRequest{SlavesUUID: allHostUUIDList}
	var res types.APISlaveProfileListResponse
	err = util.HttpPost(hostProfileListPath, &req, &res)
	if err != nil {
		return nil, nil, nil, err
	}

	for _, slave := range res.Slaves {
		hostProfile = append(hostProfile, slave)
		host = append(host, uuid2host[slave.SlaveUUId])
		delete(uuid2host, slave.SlaveUUId)
	}

	for _, lostHost := range uuid2host {
		lost = append(lost, lostHost)
	}
	return
}

func (hd *HostDaemon) UpdateHostAddToken() (token string, code int) {
	hostAddTokenUpdatePath := setting.UDCSetting.URL + "/slave/update_add_token"
	var res types.APISlaveAddToken
	err := util.HttpPost(hostAddTokenUpdatePath, nil, &res)
	if err != nil {
		log.Println(err.Error())
		return "", e.ERROR
	}

	return res.Token, e.SUCCESS
}

func (hd *HostDaemon) GetHostAddToken() (token string, code int) {
	hostAddTokenUpdatePath := setting.UDCSetting.URL + "/slave/add_token"
	var res types.APISlaveAddToken
	err := util.HttpPost(hostAddTokenUpdatePath, nil, &res)
	if err != nil {
		return "", e.ERROR
	}

	return res.Token, e.SUCCESS
}
