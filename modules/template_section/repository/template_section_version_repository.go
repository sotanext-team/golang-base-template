package repository

import (
	"context"

	"app-api/ent"
	"app-api/modules/template_section/types"
)

type TemplateSectionVersionRepository interface {
	FindLastVersion(client *ent.Client, ctx context.Context, templateID uint64) (*ent.TemplateSectionVersion, error)
	ListByThemeTemplateIDWithCursorPagination(client *ent.Client, ctx context.Context, themeTemplateID uint64, params types.TemplateSectionVersionGraphListInput) (*ent.TemplateSectionVersionConnection, error)
	Create(client *ent.Client, ctx context.Context, version *ent.TemplateSectionVersion) (*ent.TemplateSectionVersion, error)
}
