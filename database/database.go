package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	DBCon, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/booking?charset=utf8&parseTime=true")

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err.Error(),
			"errorCode": 500,
		}).Panic("Failed To Connect Database")
	}
}
