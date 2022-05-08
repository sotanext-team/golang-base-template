package db

import (
	"context"
	"database/sql"
	"fmt"

	"app-api/configs"
	"app-api/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

var client *ent.Client

// Open new DB connection
func newClient(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		panic(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func GetClient() *ent.Client {
	if client == nil {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			configs.Database.Host,
			configs.Database.Username,
			configs.Database.Password,
			configs.Database.Name,
			configs.Database.Port,
		)
		client = newClient(dsn)
		if err := client.Schema.Create(context.Background()); err != nil {
			logrus.Fatal("opening ent client", err)
		}
		logrus.Info("Database connected successfully")
	}
	return client
}
