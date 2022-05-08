package repository

import (
	"context"

	"app-api/ent"
	"app-api/modules/theme/types"
)

type ThemeRepository interface {
	FindByID(client *ent.Client, ctx context.Context, currentShop *ent.Shop, id uint64) (*ent.Theme, error)
	ListByShopID(
		client *ent.Client,
		ctx context.Context,
		currentShop *ent.Shop,
		params types.ThemeGraphListInput,
	) (*ent.ThemeConnection, error)
	Create(client *ent.Client, ctx context.Context, theme *ent.Theme) (*ent.Theme, error)
	Save(client *ent.Client, ctx context.Context, theme *ent.Theme) (*ent.Theme, error)
	Delete(client *ent.Client, ctx context.Context, id uint64) error
}
