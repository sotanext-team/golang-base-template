package db

import (
	"golang-base/configs"
	"golang-base/models"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {

	if configs.IsDev() {
		db = GetClient().Debug()
	} else {
		db = GetClient()
	}

	logrus.Info("Start migrating with ent schema")
	// Run the auto migration tool.
	err = db.AutoMigrate(&models.Shop{})
	if err != nil {
		log.Fatal("Unable to Auto migrate. Message: ", err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
