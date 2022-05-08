package seeding

import (
	"app-api/ent"
)

type Seed struct {
	Name string
	Run  func(*ent.Client) error
}

func All() []Seed {
	return []Seed{
		{
			Name: "CreateShop",
			Run: func(client *ent.Client) error {
				return createShop(client)
			},
		},
	}
}

// func bulkInsert(db *gorm.DB, fakeData []interface{}, T any) error {
// 	err := db.Model(&T).Create(&fakeData).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
