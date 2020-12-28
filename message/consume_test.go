package message

import (
	"testing"
)

func TestConnectConsumer(t *testing.T) {
	t.SkipNow()
	//mockConn, err := amqptest.Dial("amqp://localhost:5672/%2f") // will fail,
	//
	//if err == nil {
	//	t.Error("This shall fail, because no fake amqp server is running...")
	//}
	//
	//fakeServer := server.NewServer("amqp://localhost:5672/%2f")
	//fakeServer.Start()
	//
	//mockConn, err = amqptest.Dial("amqp://localhost:5672/%2f") // now it works =D
	//
	//if err != nil {
	//	t.Error(err)
	//}
	//
	////Now you can use mockConn as a real amqp connection.
	//_, err = mockConn.Channel()

	//connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	//channel, err := connection.Channel()
	//Channel = channel
	//
	//if err == nil {
	//	t.Error("RabbitMQ Connection Failed")
	//}

	ConnectConsumer()
}
