package repository

import (
	"context"

	"app-api/ent"
)

type GlobalTemplateRepository interface {
	FindByThemeTemplateId(client *ent.Client, ctx context.Context, currentShop *ent.Shop, themeTemplateId uint64) (*ent.GlobalTemplate, error)
	Create(client *ent.Client, ctx context.Context, globalTemplate *ent.GlobalTemplate) (*ent.GlobalTemplate, error)
}
