package repository

import (
	"context"

	"app-api/ent"
	"app-api/ent/globaltemplate"
)

type globalTemplateImpl struct {
}

func NewGlobalTemplateRepository() GlobalTemplateRepository {
	return &globalTemplateImpl{}
}

func (instance *globalTemplateImpl) FindByThemeTemplateId(client *ent.Client, ctx context.Context, currentShop *ent.Shop, themeTemplateId uint64) (*ent.GlobalTemplate, error) {
	return client.GlobalTemplate.Query().
		Where(
			globaltemplate.ShopIDEQ(currentShop.ID),
			// globaltemplate.ThemeTemplateIDEQ(themeTemplateId),
		).First(ctx)
}

func (instance *globalTemplateImpl) Create(client *ent.Client, ctx context.Context, globalTemplate *ent.GlobalTemplate) (*ent.GlobalTemplate, error) {
	return client.GlobalTemplate.Create().
		SetName(globalTemplate.Name).
		SetShopID(uint64(globalTemplate.ShopID)).
		SetViewCount(globalTemplate.ViewCount).
		SetInstallCount(globalTemplate.InstallCount).
		Save(ctx)
}
