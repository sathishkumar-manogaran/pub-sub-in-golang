package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	DBCon, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/booking?charset=utf8&parseTime=true")

	if err != nil {

		panic("failed to connect database")
	}
}
