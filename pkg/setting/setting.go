package setting

import (
	"log"
	"strconv"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
}

var AppSetting = &App{}

type Server struct {
	RunMode        string
	Domain         string
	HttpPort       int
	CallbackPrefix string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

var ServerSetting = &Server{}

type Mysql struct {
	User     string
	Password string
	Host     string
	Name     string
}

var MysqlSetting = &Mysql{}

type UDC struct {
	URL string
}

var UDCSetting = &UDC{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)

	mapTo("server", ServerSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	ServerSetting.CallbackPrefix = ServerSetting.Domain + ":" + strconv.Itoa(ServerSetting.HttpPort) + "/callback"

	mapTo("mysql", MysqlSetting)

	mapTo("udc", UDCSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
