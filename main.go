package main

import (
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/database"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/message"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/migrations"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

func main() {

	database.InitDB()
	migrations.Migrate()

	message.InitChannel()
	message.Publisher()
	message.Consume()

	//defer database.DBCon.Close()
	defer message.Channel.Close()
}
