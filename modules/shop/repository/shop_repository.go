package repository

import (
	"context"

	"app-api/ent"
	"app-api/models"
)

type shopImpl struct {
	dbEnt *ent.Client
}

func NewShopRepository(dbEnt *ent.Client) ShopRepository {
	return &shopImpl{
		dbEnt: dbEnt,
	}
}

type ShopRepository interface {
	CreateShop(ctx context.Context, shop models.Shop) (models.Shop, error)
	CreateShopRevert(ctx context.Context, shop models.Shop) error
	GetShop(ctx context.Context, id uint64) (*ent.Shop, error)
	GetShopByDomainE(ctx context.Context, domain string) (*ent.Shop, error)
}
