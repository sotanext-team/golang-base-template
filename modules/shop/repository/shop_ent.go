package repository

import (
	"context"

	"app-api/ent"
	"app-api/ent/shop"
	"app-api/models"

	"github.com/sirupsen/logrus"
)

func (instance *shopImpl) CreateShop(ctx context.Context, shop models.Shop) (models.Shop, error) {
	shopInput := ent.CreateShopInput{
		DefaultDomain: shop.DefaultDomain,
		CustomDomain:  shop.CustomDomain,
	}
	logrus.Info(shop.DefaultDomain)
	logrus.Info(shop.CustomDomain)
	shopCreate, err := instance.dbEnt.Shop.Create().SetInput(shopInput).Save(ctx)
	if err != nil {
		logrus.Info(err)
		return shop, err
	}
	shop.CreatedAt = shopCreate.CreatedAt
	shop.UpdatedAt = shopCreate.UpdatedAt
	shop.ID = shopCreate.ID
	return shop, err
}

func (instance *shopImpl) CreateShopRevert(ctx context.Context, s models.Shop) error {
	client := instance.dbEnt
	logrus.Info(s.DefaultDomain)
	logrus.Info(s.CustomDomain)
	_, err := client.Shop.Delete().Where(shop.DefaultDomainEQ(s.DefaultDomain)).Exec(ctx)
	return err
}

func (instance *shopImpl) GetShop(ctx context.Context, id uint64) (*ent.Shop, error) {
	client := instance.dbEnt
	return client.Shop.Get(ctx, id)
}

func (instance *shopImpl) GetShopByDomainE(ctx context.Context, domain string) (*ent.Shop, error) {
	client := instance.dbEnt
	return client.Shop.Query().Where(shop.DefaultDomainEQ(domain)).First(ctx)
}
