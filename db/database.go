package db

import (
	"context"

	"app-api/configs"
	"app-api/ent"

	"github.com/sirupsen/logrus"
)

var entDb *ent.Client

func InitDB() {

	if configs.IsDev() {
		entDb = GetClient().Debug()
	} else {
		entDb = GetClient()
	}

	logrus.Info("Start migrating with ent schema")
	// Run the auto migration tool.
	if err := entDb.Schema.Create(context.Background()); err != nil {
		logrus.Fatalf("failed creating schema resources: %v", err)
	}

	// err = db.AutoMigrate(
	// 	&models.User{}, &models.UserMeta{},
	// 	&models.Shop{}, &models.UserShop{}, &models.ShopMeta{}, &models.ShopEvent{},
	// 	&models.GlobalStyle{}, &models.Widget{},
	// 	&models.Theme{}, &models.ThemePage{}, &models.ThemeGlobalStyle{}, &models.ThemeMeta{},
	// 	&models.ThemeTemplate{},
	// 	&models.GlobalTemplate{},
	// 	&models.TemplateSection{}, &models.BkTemplateSection{}, &models.TemplateSectionVersion{},
	// )
	// if err != nil {
	// 	log.Fatal("Unable to Auto migrate. Message: ", err.Error())
	// }
}

func GetEntDB() *ent.Client {
	return entDb
}
