package models

type Offers struct {
	Offer []Offer `json:"offers"`
}

type Offer struct {
	CmOfferId    string `json:"cm_offer_id"`
	Hotel        `json:"hotel"`
	Room         `json:"room"`
	RatePlan     `json:"rate_plan"`
	OriginalData `json:"original_data"`
	Capacity     `json:"capacity"`
	Number       `json:"number"`
	Price        `json:"price"`
	Currency     `json:"currency"`
	CheckIn      `json:"check_in"`
	CheckOut     `json:"check_out"`
	Fees         []Fee `json:"fees"`
}

type Hotel struct {
	HotelId   string  `json:"hotel_id" gorm:"primary_key;size:10"`
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Telephone string  `json:"telephone"`
	//Amenities   []string `json:"amenities" gorm:"-"`
	Description string     `json:"description"`
	RoomCount   int8       `json:"room_count"`
	Currency    string     `json:"currency"`
	Rooms       []Room     `gorm:"foreignKey:hotel_id"`
	RatePlans   []RatePlan `gorm:"foreignKey:hotel_id"`
}

type Room struct {
	RoomId      string `json:"room_id" gorm:"primary_key"`
	HotelId     string `json:"hotel_id" gorm:"size:10"`
	Description string `json:"description"`
	Name        string `json:"name"`
	//Capacity    `json:"capacity" gorm:"-"`
}

type RatePlan struct {
	RatePlanId string `json:"rate_plan_id" gorm:"primary_key"`
	HotelId    string `json:"hotel_id" gorm:"size:10"`
	//CancellationPolicy []CancellationPolicy `json:"cancellation_policy" gorm:"-"`
	Name string `json:"name"`
	//OtherConditions    []string             `json:"other_conditions"`
	MealPlan string `json:"meal_plan"`
}

type Capacity struct {
	MaxAdults     int8 `json:"max_adults"`
	ExtraChildren int8 `json:"extra_children"`
}

type CancellationPolicy struct {
	Type              string `json:"type"`
	ExpiresDaysBefore int8   `json:"expires_days_before"`
}

type OriginalData struct {
	GuaranteePolicy `json:"GuaranteePolicy"`
}

type GuaranteePolicy struct {
	Required bool `json:"Required"`
}

type Number int
type Price int
type Currency string
type CheckIn string
type CheckOut string

type Fee struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Included    bool    `json:"included"`
	Percent     float32 `json:"percent"`
}
