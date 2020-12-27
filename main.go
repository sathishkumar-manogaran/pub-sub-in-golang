package main

import (
	"fmt"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/database"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/migrations"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/publisher"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/subscriber"
)

func main() {
	fmt.Println("Initializer")

	/*offers := models.Offers{Offer: []models.Offer{{
		CmOfferId: "8f6995366e854c9faf1d9f3d233702b8",
		Hotel: models.Hotel{
			HotelId:   "BH~46456",
			Name:      "Hawthorn Suites by Wyndham Eagle CO",
			Country:   "US",
			Address:   "0315 Chambers Avenue, 81631",
			Latitude:  39.660193,
			Longitude: -106.824123,
			Telephone: "+1-970-3283000",
			Amenities: []string{"Business Centre",
				"Fitness Room/Gym",
				"Pet Friendly",
				"Disabled Access",
				"Air Conditioned",
				"Free WIFI",
				"Elevator / Lift",
				"Parking"},
			Description: "Stay a while in beautiful mountain country at this Hawthorn Suites by Wyndham Eagle CO hotel, just off Interstate 70, only 6 miles from the Vail/Eagle Airport and close to skiing, golfing, Eagle River and great restaurants. Pets are welcome at this hotel",
			RoomCount:   1,
			Currency:    "USD",
		},
		Room: models.Room{
			HotelId:     "BH~46456",
			RoomId:      "S2Q",
			Description: "JUNIOR SUITES WITH 2 QUEEN BEDS",
			Name:        "JUNIOR SUITES WITH 2 QUEEN BEDS",
			Capacity: models.Capacity{
				MaxAdults:     2,
				ExtraChildren: 2,
			},
		},
		RatePlan: models.RatePlan{
			HotelId:    "BH~46456",
			RatePlanId: "BAR",
			CancellationPolicy: []models.CancellationPolicy{{
				Type:              "Free cancellation",
				ExpiresDaysBefore: 2,
			}},
			Name: "BEST AVAILABLE RATE",
			OtherConditions: []string{"CXL BY 2 DAYS PRIOR TO ARRIVAL-FEE 1 NIGHT 2 DAYS PRIOR TO ARRIVAL",
				"BEST AVAILABLE RATE"},
			MealPlan: "Room only",
		},
		OriginalData: models.OriginalData{GuaranteePolicy: models.GuaranteePolicy{Required: true}},
		Capacity: models.Capacity{
			MaxAdults:     2,
			ExtraChildren: 2,
		},
		Number:   1,
		Price:    1520,
		Currency: "USD",
		CheckIn:  "2020-11-18",
		CheckOut: "2020-11-20",
		Fees:     []models.Fee{{Type: "CountyTax", Description: "COUNTY TAX PER STAY", Included: true, Percent: 17.5}},
	},
	},
	}*/

	////initialize the database
	database.InitDB()

	migrations.Migrate()

	publisher.Publisher()
	subscriber.Consume()

	///finally close the connection when you are done
	defer database.DBCon.Close()
}
