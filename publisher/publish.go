package publisher

import (
	"github.com/streadway/amqp"
	"os"
)

func Publisher() {
	// Get the connection string from the environment variable
	url := os.Getenv("AMQP_URL")

	//If it doesnt exist, use the default connection string
	if url == "" {
		url = "amqp://guest:guest@localhost:5672"
	}

	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	// Create a channel from the connection. We'll use channels to access the data in the queue rather than the
	// connection itself
	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	// We create an exahange that will bind to the queue to send and receive messages
	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	publishMessage := `{
  "offers": [
    {
      "cm_offer_id": "8f6995366e854c9faf1d9f3d233702b8",
      "hotel": {
        "hotel_id": "BH~46456",
        "name": "Hawthorn Suites by Wyndham Eagle CO",
        "country": "US",
        "address": "0315 Chambers Avenue, 81631",
        "latitude": 39.660193,
        "longitude": -106.824123,
        "telephone": "+1-970-3283000",
        "amenities": [
          "Business Centre",
          "Fitness Room/Gym",
          "Pet Friendly",
          "Disabled Access",
          "Air Conditioned",
          "Free WIFI",
          "Elevator / Lift",
          "Parking"
        ],
        "description": "Stay a while in beautiful mountain country at this Hawthorn Suites by Wyndham Eagle CO hotel, just off Interstate 70, only 6 miles from the Vail/Eagle Airport and close to skiing, golfing, Eagle River and great restaurants. Pets are welcome at this h",
        "room_count": 1,
        "currency": "USD"
      },
      "room": {
        "hotel_id": "BH~46456",
        "room_id": "S2Q",
        "description": "JUNIOR SUITES WITH 2 QUEEN BEDS",
        "name": "JUNIOR SUITES WITH 2 QUEEN BEDS",
        "capacity": {
          "max_adults": 2,
          "extra_children": 2
        }
      },
      "rate_plan": {
        "hotel_id": "BH~46456",
        "rate_plan_id": "BAR",
        "cancellation_policy": [
          {
            "type": "Free cancellation",
            "expires_days_before": 2
          }
        ],
        "name": "BEST AVAILABLE RATE",
        "other_conditions": [
          "CXL BY 2 DAYS PRIOR TO ARRIVAL-FEE 1 NIGHT 2 DAYS PRIOR TO ARRIVAL",
          "BEST AVAILABLE RATE"
        ],
        "meal_plan": "Room only"
      },
      "original_data": {
        "GuaranteePolicy": {
          "Required": true
        }
      },
      "capacity": {
        "max_adults": 2,
        "extra_children": 2
      },
      "number": 1,
      "price": 1520,
      "currency": "USD",
      "check_in": "2020-11-18",
      "check_out": "2020-11-20",
      "fees": [
        {
          "type": "CountyTax",
          "description": "COUNTY TAX PER STAY",
          "included": true,
          "percent": 17.5
        }
      ]
    }
  ]
}`

	// We create a message to be sent to the queue.
	// It has to be an instance of the aqmp publishing struct
	message := amqp.Publishing{
		Body: []byte(publishMessage),
	}

	// We publish the message to the exahange we created earlier
	err = channel.Publish("events", "random-key", false, false, message)

	if err != nil {
		panic("error publishing a message to the queue:" + err.Error())
	}

	// We create a queue named Test
	_, err = channel.QueueDeclare("test", true, false, false, false, nil)

	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind("test", "#", "events", false, nil)

	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}
}
