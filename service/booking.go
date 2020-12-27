package service

import (
	"fmt"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/database"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/models"
)

// Offers Iterate through offers and save / create values
func Offers(offers models.Offers) {
	if offers.Offer != nil {
		for _, v := range offers.Offer {
			Hotel(v.Hotel)
			Room(v.Room)
			RatePlan(v.RatePlan)
		}
	}
	//GetHotelValue()
}

func Hotel(hotel models.Hotel) {
	find := database.DBCon.Where(&hotel, hotel.HotelId)
	if 0 == find.RowsAffected {
		database.DBCon.Create(&hotel)
	} else {
		database.DBCon.Updates(&hotel)
	}
	fmt.Println(hotel)
}

func Room(room models.Room) {
	find := database.DBCon.Where(&room, room.HotelId)
	if 0 == find.RowsAffected {
		database.DBCon.Create(&room)
	} else {
		database.DBCon.Updates(&room)
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

func RatePlan(ratePlan models.RatePlan) {
	find := database.DBCon.Where(&ratePlan, ratePlan.HotelId)
	if 0 == find.RowsAffected {
		database.DBCon.Create(&ratePlan)
	} else {
		database.DBCon.Updates(&ratePlan)
	}
	fmt.Println(ratePlan)

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
