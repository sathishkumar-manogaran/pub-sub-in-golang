package migrations

import (
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/database"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/models"
)

func Migrate() {

	//DBCon.CreateTable(&Hotel{}, &RatePlan{}, &Room{})

	database.DBCon.AutoMigrate(&models.Hotel{}, &models.RatePlan{}, &models.Room{})

	database.DBCon.Model(&models.Room{}).AddForeignKey("hotel_id", "hotel(hotel_id)", "CASCADE", "RESTRICT")

}
