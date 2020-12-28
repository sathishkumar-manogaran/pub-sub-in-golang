package service

import (
	"fmt"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/database"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/models"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Get(id string) (*models.Hotel, error)
	Create(id string, name string) error
	Hotel(*models.Hotel)
}

// Offers Iterate through offers and save / create values
func Offers(offers *models.Offers) {
	if offers.Offer != nil {
		for _, v := range offers.Offer {
			Hotel(&v.Hotel)
			Room(&v.Room)
			RatePlan(&v.RatePlan)
		}
	}
	//GetHotelValue()
}

func Hotel(hotel *models.Hotel) {
	result := database.DBCon.Find(&hotel)
	if result.RowsAffected == 0 {
		database.DBCon.Save(&hotel)
	} else {
		database.DBCon.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Updates(&hotel)
	}

	fmt.Println(hotel)
}

func Room(room *models.Room) {
	result := database.DBCon.Find(&room)
	if result.RowsAffected == 0 {
		database.DBCon.Save(&room)
	} else {
		database.DBCon.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Updates(&room)
	}

	fmt.Println(room)

	/**
	Just for reference
	*/
	//var result2 models.Room
	//if err := database.DBCon.First(&result2, datatypes.JSONQuery("capacity").HasKey("max_adults")).Error; err != nil {
	//t.Fatalf("failed to find user with json key, got error %v", err)
	//}
	//fmt.Println("Room capacity:: ", result2.Capacity)
}

func RatePlan(ratePlan *models.RatePlan) {
	result := database.DBCon.Find(&ratePlan)
	if result.RowsAffected == 0 {
		database.DBCon.Save(&ratePlan)
	} else {
		database.DBCon.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Updates(&ratePlan)
	}

	/**
	Just for reference
	*/
	//var result2 models.RatePlan
	//if err := database.DBCon.First(&result2, datatypes.JSONQuery("cancellation_policy").HasKey("type", "type")).Error; err != nil {
	//	//t.Fatalf("failed to find user with json key, got error %v", err)
	//}
}

// GetHotelValue Just for reference
func GetHotelValue() {

	hotel := models.Hotel{}
	database.DBCon.Find(&hotel).Scan(hotel)
	fmt.Println("Final Hotel Value :: ")
	//amenities, _ := hotel
	fmt.Println(hotel)
}
