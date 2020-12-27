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

	// We consume data in the queue named test using the channel we created in go.
	msgs, err := channel.Consume("test", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	// We loop through the messages in the queue and print them to the console.
	// The msgs will be a go channel, not an amqp channel
	for msg := range msgs {
		//print the message to the console
		//fmt.Println("message received: " + string(msg.Body))

		j := msg.Body

		var offers models.Offers

		err := json.Unmarshal(j, &offers)

		if err != nil {
			log.Fatalf("Unable to decode the json")
		}

		fmt.Println(offers.Offer[0].Amenities)

		// Acknowledge that we have received the message so it can be removed from the queue
		msg.Ack(true)
	}

	// We close the connection after the operation has completed.
	defer connection.Close()
}
