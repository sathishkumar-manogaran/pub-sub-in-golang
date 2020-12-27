package message

import (
	"encoding/json"
	"fmt"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/models"
	log "github.com/sirupsen/logrus"
)

func Consume() {

	messages, err := Channel.Consume("test", "", false, false, false, false, nil)

	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err.Error(),
			"errorCode": 500,
		}).Panic("error consuming the queue")
	}

	for msg := range messages {

		j := msg.Body

		var offers models.Offers

		err := json.Unmarshal(j, &offers)

		if err != nil {
			log.WithFields(log.Fields{
				"msg":       err.Error(),
				"errorCode": 500,
			}).Fatal("Unable to decode the json")
		}

		fmt.Println(offers.Offer[0].Amenities)

		msg.Ack(true)
	}
}
