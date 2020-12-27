package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	//DBCon, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/booking?charset=utf8&parseTime=true")
	dsn := "root:root@tcp(127.0.0.1:3306)/booking?charset=utf8mb4&parseTime=True&loc=Local"
	DBCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err.Error(),
			"errorCode": 500,
		}).Panic("Failed To Connect Database")
	}
}
