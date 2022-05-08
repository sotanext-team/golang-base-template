package repository

import (
	"context"

	"golang-base/models"
)

func (instance *shopImpl) CreateShop(ctx context.Context, shop models.Shop) (*models.Shop, error) {
	client := instance.db
	err := client.Save(&shop).Error
	if err != nil {
		return nil, err
	}
	return &shop, err
}

func (instance *shopImpl) GetShop(ctx context.Context, id uint64) (*models.Shop, error) {
	client := instance.db
	shop := models.Shop{}
	err := client.Where("id = ?", id).First(&shop).Error
	if err != nil {
		return nil, err
	}
	return &shop, err
}

func (instance *shopImpl) GetShops(ctx context.Context) ([]models.Shop, error) {
	client := instance.db
	shops := []models.Shop{}
	err := client.Find(&shops).Error
	if err != nil {
		return nil, err
	}
	return shops, err
}
