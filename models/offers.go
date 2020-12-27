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
	HotelId     string   `json:"hotel_id,primary_key"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Address     string   `json:"address"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Telephone   string   `json:"telephone"`
	Amenities   []string `json:"amenities,blob"`
	Description string   `json:"description"`
	RoomCount   int8     `json:"room_count"`
	Currency    string   `json:"currency"`
}

type Room struct {
	HotelId     string `json:"hotel_id"`
	RoomId      string `json:"room_id,primary_key"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Capacity    `json:"capacity"`
}

type Capacity struct {
	MaxAdults     int8 `json:"max_adults"`
	ExtraChildren int8 `json:"extra_children"`
}

type RatePlan struct {
	HotelId            string               `json:"hotel_id"`
	RatePlanId         string               `json:"rate_plan_id,primary_key"`
	CancellationPolicy []CancellationPolicy `json:"cancellation_policy"`
	Name               string               `json:"name"`
	OtherConditions    []string             `json:"other_conditions"`
	MealPlan           string               `json:"meal_plan"`
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
