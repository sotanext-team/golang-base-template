package seeding

import (
	"context"

	"app-api/ent"

	"github.com/bxcodec/faker/v3"

	"github.com/mitchellh/mapstructure"
)

func createShop(client *ent.Client) error {
	fakeData := make([]interface{}, 0)
	numberOfFakeData := 10
	for i := 0; i < numberOfFakeData; i++ {
		p := struct {
			DefaultDomain string `json:"defaultDomain" gorm:"size:50" faker:"username"`
			CustomDomain  string `json:"customDomain" gorm:"size:50" faker:"username"`
		}{}
		_ = faker.FakeData(&p)

		p.DefaultDomain = p.DefaultDomain + ".eshs.com"
		p.CustomDomain = p.CustomDomain + ".eshs.com"

		fakeData = append(fakeData, p)
	}

	shops := []ent.Shop{}
	mapstructure.Decode(fakeData, &shops)
	shops = append(shops, ent.Shop{
		DefaultDomain: "noragem.eshs.com",
		CustomDomain:  "noragem.eshs.com",
	})

	bulk := make([]*ent.ShopCreate, len(shops))

	for i, shop := range shops {
		bulk[i] = client.Shop.Create().SetDefaultDomain(shop.DefaultDomain).SetCustomDomain(shop.CustomDomain)
	}

	ctx := context.Background()

	_, err := client.Shop.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
