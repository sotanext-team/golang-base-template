package repository

import (
	"context"

	"golang-base/models"

	"gorm.io/gorm"
)

type shopImpl struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopImpl{
		db: db,
	}
}

type ShopRepository interface {
	CreateShop(ctx context.Context, shop models.Shop) (*models.Shop, error)
	GetShop(ctx context.Context, id uint64) (*models.Shop, error)
	GetShops(ctx context.Context) ([]models.Shop, error)
}
