package models

import (
	"fmt"
	"github.com/PenguinCats/unison-web-backend/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
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

	if setting.ServerSetting.RunMode == "debug" {
		newLoagger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			})

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLoagger,
		})
	} else {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}

func NewContextForTransaction() *gorm.DB {
	return db.Begin()
}
