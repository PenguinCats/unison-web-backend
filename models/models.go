package models

import (
	"fmt"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlSetting.User,
		setting.MysqlSetting.Password,
		setting.MysqlSetting.Host,
		setting.MysqlSetting.Name,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}
