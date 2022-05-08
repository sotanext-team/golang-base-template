package db

import (
	"fmt"

	"golang-base/configs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var client *gorm.DB
var err error

// Open new DB connection
func newClient(databaseUrl string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
}

func GetClient() *gorm.DB {
	if client == nil {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			configs.Database.Host,
			configs.Database.Username,
			configs.Database.Password,
			configs.Database.Name,
			configs.Database.Port,
		)
		client, err = newClient(dsn)
		if err != nil {
			panic(err)
		}
		logrus.Info("Database connected successfully")
	}
	return client
}
