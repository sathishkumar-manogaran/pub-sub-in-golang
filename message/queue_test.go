package message

import (
	"github.com/streadway/amqp"
	"testing"
)

func TestConnectToQueue(t *testing.T) {
	queue, err := connectToQueue()
	if queue == nil {
		t.Error("Queue Failed")
	}
	if err != nil {
		t.Error("Error Occurred", err)
	}
}

func TestGetChannel(t *testing.T) {
	t.SkipNow()
	var err error
	var connection *amqp.Connection

	getChannel(err, connection)
}
