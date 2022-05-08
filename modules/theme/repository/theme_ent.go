package repository

import (
	"context"
	"fmt"

	"app-api/ent"
	"app-api/ent/theme"
	"app-api/modules/theme/types"
)

type themeImpl struct {
}

func NewThemePostgres() ThemeRepository {
	return &themeImpl{}
}

func (instance *themeImpl) Create(client *ent.Client, ctx context.Context, theme *ent.Theme) (*ent.Theme, error) {
	theme, err := client.Theme.
		Create().
		SetName(theme.Name).
		SetThumbnail(theme.Thumbnail).
		SetPublish(theme.Publish).
		SetShopID(theme.ShopID).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating theme: %w", err)
	}
	return theme, nil
}

func (instance *themeImpl) FindByID(client *ent.Client, ctx context.Context, currentShop *ent.Shop, id uint64) (*ent.Theme, error) {
	theme, err := client.Theme.Query().Where(theme.ShopIDEQ(currentShop.ID), theme.ID(id)).Only(ctx)

	if err != nil {
		return nil, err
	}
	return theme, nil
}

func (instance *themeImpl) Save(client *ent.Client, ctx context.Context, theme *ent.Theme) (*ent.Theme, error) {
	return client.Theme.
		UpdateOneID(theme.ID).
		SetName(theme.Name).
		SetThumbnail(theme.Thumbnail).
		SetPublish(theme.Publish).
		Save(ctx)
}

func (instance *themeImpl) Delete(client *ent.Client, ctx context.Context, id uint64) error {
	return client.Theme.DeleteOneID(id).Exec(ctx)
}

func (instance *themeImpl) ListByShopID(
	client *ent.Client,
	ctx context.Context,
	currentShop *ent.Shop,
	params types.ThemeGraphListInput,
) (*ent.ThemeConnection, error) {
	return client.Theme.Query().
		Where(theme.ShopIDEQ(currentShop.ID)).
		Paginate(ctx, params.After, params.First, params.Before, params.Last,
			ent.WithThemeOrder(params.OrderBy),
		)
}
