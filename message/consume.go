package message

import (
	"encoding/json"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/models"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/service"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func ConnectConsumer() {
	messages, err := Channel.Consume("test", "", false, false, false, false, nil)

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err.Error(),
			"errorCode": 500,
		}).Panic("error consuming the queue")
	}

	startConsumer(messages)
}

func startConsumer(messages <-chan amqp.Delivery) {
	for message := range messages {
		requestBody := message.Body

		log.WithFields(log.Fields{
			"message": requestBody,
		}).Info("Raw Request Body")

		var offers models.Offers
		err := json.Unmarshal(requestBody, &offers)
		if err != nil {
			log.WithFields(log.Fields{
				"message":   err.Error(),
				"errorCode": 500,
			}).Fatal("Unable to decode the json")
		}

		err = message.Ack(true)
		if err != nil {
			log.WithFields(log.Fields{
				"message":   err.Error(),
				"errorCode": 500,
			}).Fatal("Unable to message acknowledge")
		}

		// process to service layer and save to db
		service.Offers(&offers)
	}
}
