package repository

import (
	"context"

	"app-api/ent"
	"app-api/ent/themetemplate"
	"app-api/modules/theme_template/types"
)

type ThemeTemplateRepository interface {
	ListByThemeIDAndPageTypeWithCursorPagination(
		client *ent.Client, ctx context.Context,
		currentShop *ent.Shop,
		themeId uint,
		pageType themetemplate.PageType,
		params types.ThemeTemplateGraphListInput,
	) (*ent.ThemeTemplateConnection, error)
	ListByThemeIDAndPageTypeWithCursorPaginationIntrash(
		client *ent.Client, ctx context.Context,
		currentShop *ent.Shop,
		themeId uint,
		pageType themetemplate.PageType,
		params types.ThemeTemplateGraphListInput,
	) (*ent.ThemeTemplateConnection, error)
	FindByName(
		client *ent.Client, ctx context.Context,
		currentShop *ent.Shop,
		name string,
		themeId uint64,
		pageType themetemplate.PageType,
	) (*ent.ThemeTemplate, error)
	FindByID(client *ent.Client, ctx context.Context, currentShop *ent.Shop, id uint64) (*ent.ThemeTemplate, error)
	FindByIDUnscoped(client *ent.Client, ctx context.Context, currentShop *ent.Shop, id uint64) (*ent.ThemeTemplate, error)
	Create(client *ent.Client, ctx context.Context, themeTemplate *ent.ThemeTemplate) (*ent.ThemeTemplate, error)
	Save(client *ent.Client, ctx context.Context, themeTemplate *ent.ThemeTemplate) (*ent.ThemeTemplate, error)
	Restore(client *ent.Client, ctx context.Context, id uint64) (*ent.ThemeTemplate, error)
	Delete(client *ent.Client, ctx context.Context, id uint64) error
	ForceDelete(client *ent.Client, ctx context.Context, id uint64) error
	SetDefault(client *ent.Client, ctx context.Context, currentShop *ent.Shop, themeId uint64, value bool) error
	// BatchUpdates(client *ent.Client, ctx context.Context, themeId uint, mapUpdates map[string]interface{}) error
}
