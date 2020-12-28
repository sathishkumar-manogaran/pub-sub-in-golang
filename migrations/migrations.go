package migrations

import (
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/database"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/models"
	log "github.com/sirupsen/logrus"
)

func Migrate() {

	err := database.DBCon.AutoMigrate(
		&models.Hotel{},
		&models.Room{},
		&models.RatePlan{},
	)
	if err != nil {
		log.WithFields(log.Fields{
			"msg":       err.Error(),
			"errorCode": 500,
		}).Error("error creating tables")
	}

}
