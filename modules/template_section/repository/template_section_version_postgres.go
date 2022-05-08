package repository

import (
	"context"

	"app-api/ent"
	"app-api/ent/templatesectionversion"
	"app-api/modules/template_section/types"
)

type templateSectionVersionImpl struct {
}

func NewTemplateSectionVersionRepository() TemplateSectionVersionRepository {
	return &templateSectionVersionImpl{}
}

func (instance *templateSectionVersionImpl) FindLastVersion(client *ent.Client, ctx context.Context, templateID uint64) (*ent.TemplateSectionVersion, error) {
	return client.TemplateSectionVersion.Query().
		Where(
			templatesectionversion.ThemeTemplateIDEQ(templateID),
		).
		Order(
			ent.Desc(templatesectionversion.FieldID),
		).
		First(ctx)
}

func (instance *templateSectionVersionImpl) Create(client *ent.Client, ctx context.Context, version *ent.TemplateSectionVersion) (*ent.TemplateSectionVersion, error) {
	return client.TemplateSectionVersion.Create().
		SetName(version.Name).
		SetVersion(version.Version).
		SetThemeTemplateID(version.ThemeTemplateID).
		Save(ctx)
}

func (instance *templateSectionVersionImpl) ListByThemeTemplateIDWithCursorPagination(client *ent.Client, ctx context.Context, themeTemplateID uint64, params types.TemplateSectionVersionGraphListInput) (*ent.TemplateSectionVersionConnection, error) {
	return client.TemplateSectionVersion.Query().
		Where(templatesectionversion.ThemeTemplateIDEQ(themeTemplateID)).
		Paginate(ctx, params.After, params.First, params.Before, params.Last,
			ent.WithTemplateSectionVersionOrder(params.OrderBy),
		)
}
