package repository

import (
	"context"
	"time"

	"app-api/ent"
	"app-api/ent/themetemplate"
	"app-api/modules/theme_template/types"
)

type themeTemplateImpl struct {
}

func NewThemeTemplateRepository() ThemeTemplateRepository {
	return &themeTemplateImpl{}
}

func (instance *themeTemplateImpl) ListByThemeIDAndPageTypeWithCursorPagination(
	client *ent.Client, ctx context.Context,
	currentShop *ent.Shop,
	themeId uint,
	pageType themetemplate.PageType,
	params types.ThemeTemplateGraphListInput,
) (*ent.ThemeTemplateConnection, error) {
	return client.ThemeTemplate.Query().
		Where(
			themetemplate.ThemeIDEQ(uint64(themeId)),
			themetemplate.ShopIDEQ(currentShop.ID),
			themetemplate.DeletedAtIsNil(),
			themetemplate.PageTypeEQ(pageType),
		).
		Paginate(ctx, params.After, params.First, params.Before, params.Last,
			ent.WithThemeTemplateOrder(params.OrderBy),
		)
}

func (instance *themeTemplateImpl) ListByThemeIDAndPageTypeWithCursorPaginationIntrash(
	client *ent.Client, ctx context.Context,
	currentShop *ent.Shop,
	themeId uint,
	pageType themetemplate.PageType,
	params types.ThemeTemplateGraphListInput,
) (*ent.ThemeTemplateConnection, error) {
	return client.ThemeTemplate.Query().
		Where(
			themetemplate.ThemeIDEQ(uint64(themeId)),
			themetemplate.ShopIDEQ(currentShop.ID),
			themetemplate.DeletedAtNotNil(),
			themetemplate.PageTypeEQ(pageType),
		).
		Paginate(ctx, params.After, params.First, params.Before, params.Last,
			ent.WithThemeTemplateOrder(params.OrderBy),
		)
}

func (instance *themeTemplateImpl) FindByName(client *ent.Client, ctx context.Context, currentShop *ent.Shop, name string, themeId uint64, pageType themetemplate.PageType) (*ent.ThemeTemplate, error) {
	return client.Debug().ThemeTemplate.Query().
		Where(
			themetemplate.ShopIDEQ(currentShop.ID),
			themetemplate.ThemeIDEQ(themeId),
			themetemplate.NameEQ(name),
			themetemplate.PageTypeEQ(pageType),
			themetemplate.DeletedAtIsNil(),
		).
		First(ctx)
}

func (instance *themeTemplateImpl) Create(client *ent.Client, ctx context.Context, themeTemplate *ent.ThemeTemplate) (*ent.ThemeTemplate, error) {
	return client.ThemeTemplate.Create().
		SetShopID(themeTemplate.ShopID).
		SetThemeID(themeTemplate.ThemeID).
		SetName(themeTemplate.Name).
		SetPageType(themeTemplate.PageType).
		SetDefault(themeTemplate.Default).
		Save(ctx)
}

func (instance *themeTemplateImpl) FindByID(client *ent.Client, ctx context.Context, currentShop *ent.Shop, id uint64) (*ent.ThemeTemplate, error) {
	return client.ThemeTemplate.Query().
		Where(
			themetemplate.ShopIDEQ(currentShop.ID),
			themetemplate.ID(id),
			themetemplate.DeletedAtIsNil(),
		).Only(ctx)
}

func (instance *themeTemplateImpl) FindByIDUnscoped(client *ent.Client, ctx context.Context, currentShop *ent.Shop, id uint64) (*ent.ThemeTemplate, error) {
	return client.ThemeTemplate.Query().
		Where(
			themetemplate.ShopIDEQ(currentShop.ID),
			themetemplate.ID(id),
		).Only(ctx)
}

func (instance *themeTemplateImpl) Save(client *ent.Client, ctx context.Context, themeTemplate *ent.ThemeTemplate) (*ent.ThemeTemplate, error) {
	return client.ThemeTemplate.
		UpdateOneID(themeTemplate.ID).
		SetName(themeTemplate.Name).
		SetDefault(themeTemplate.Default).
		Save(ctx)
}

func (instance *themeTemplateImpl) Delete(client *ent.Client, ctx context.Context, id uint64) error {
	_, err := client.ThemeTemplate.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return err
}

func (instance *themeTemplateImpl) ForceDelete(client *ent.Client, ctx context.Context, id uint64) error {
	return client.ThemeTemplate.DeleteOneID(id).Exec(ctx)
}

func (instance *themeTemplateImpl) SetDefault(client *ent.Client, ctx context.Context, currentShop *ent.Shop, themeId uint64, value bool) error {
	_, err := client.ThemeTemplate.
		Update().
		Where(
			themetemplate.ShopIDEQ(currentShop.ID),
		).
		SetDefault(value).
		Save(ctx)
	return err
}

func (instance *themeTemplateImpl) Restore(client *ent.Client, ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	return client.ThemeTemplate.
		UpdateOneID(id).
		ClearDeletedAt().
		Save(ctx)
}

// func (instance *themeTemplateImpl) BatchUpdates( themeId uint, mapUpdates map[string]interface{}) error {
// 	// Update with map
// 	return db.Model(models.ThemeTemplate{}).Where("theme_id = ?", themeId).Updates(mapUpdates).Error
// }
