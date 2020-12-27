package subscriber

import (
	"encoding/json"
	"fmt"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/models"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func Consume() {

	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://guest:guest@localhost:5672"
	}

	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()

	msgs, err := channel.Consume("test", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	for msg := range msgs {

		j := msg.Body

		var offers models.Offers

		err := json.Unmarshal(j, &offers)

		if err != nil {
			log.Fatalf("Unable to decode the json")
		}

		fmt.Println(offers.Offer[0].Amenities)

		msg.Ack(true)
	}

	defer connection.Close()
}
