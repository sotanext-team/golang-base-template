package usecase

import (
	"context"

	"golang-base/models"
	"golang-base/modules/shop/repository"

	"gorm.io/gorm"
)

type shopUseCase struct {
	shopRepo repository.ShopRepository
	database *gorm.DB
}

func NewShopUseCase(db *gorm.DB) ShopUseCase {
	shopRepo := repository.NewShopRepository(db)
	return &shopUseCase{
		shopRepo: shopRepo,
	}
}

type ShopUseCase interface {
	CreateShop(ctx context.Context, inputShop models.Shop) (*models.Shop, error)
	GetShop(ctx context.Context, id uint64) (*models.Shop, error)
	GetShops(ctx context.Context) ([]models.Shop, error)
}

func (instance *shopUseCase) CreateShop(ctx context.Context, inputShop models.Shop) (*models.Shop, error) {
	shop, err := instance.shopRepo.CreateShop(ctx, inputShop)
	if err != nil {
		return nil, err
	}
	// Call DeployShop
	// err = grpcLibs.DeployServiceDeployShop(shop.DefaultDomain, "ap-southeast-1", "", "", "")
	if err != nil {
		return nil, err
	}
	return shop, nil
}

func (instance *shopUseCase) GetShop(ctx context.Context, id uint64) (*models.Shop, error) {
	return instance.shopRepo.GetShop(ctx, id)
}

func (instance *shopUseCase) GetShops(ctx context.Context) ([]models.Shop, error) {
	return instance.shopRepo.GetShops(ctx)
}
