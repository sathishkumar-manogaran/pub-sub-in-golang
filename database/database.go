package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	// Get the connection string from the environment variable
	url := os.Getenv("BOOKING_DB_URL")

	if url == "" {
		url = "root:root@tcp(localhost:3306)/booking"
	}
	var err error

	dsn := url + "?charset=utf8mb4&parseTime=True&loc=Local"
	DBCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err.Error(),
			"errorCode": 500,
		}).Panic("Failed To Connect Database")
	}
}
