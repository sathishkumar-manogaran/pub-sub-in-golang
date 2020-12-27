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

	//connect to postgres database
	DBCon, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/booking?charset=utf8&parseTime=true")

	//where myhost is port is the port postgres is running on
	//user is your postgres use name
	//password is your postgres password
	if err != nil {

		panic("failed to connect database")
	}
}
