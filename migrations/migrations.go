package migrations

import (
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/database"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/models"
)

func Migrate() {

	//DBCon.CreateTable(&Hotel{}, &RatePlan{}, &Room{})

	database.DBCon.AutoMigrate(&models.Hotel{}, &models.RatePlan{}, &models.Room{})
	//we take the structs we created earlier to represent tables and create tables from them.
	//for example models.Users{} will create a table called users  with the fields we defined in that struct as the table fields.
	///if by any chance you ever add another struct you need to create a table for you can add it here.

	database.DBCon.Model(&models.Room{}).AddForeignKey("hotel_id", "hotel(hotel_id)", "CASCADE", "RESTRICT")

}
