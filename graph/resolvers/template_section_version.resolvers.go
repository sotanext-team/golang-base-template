package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app-api/ent"
	"context"

	esContext "github.com/es-hs/es-helper/context"
)

func (r *queryResolver) TemplateSectionVersions(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TemplateSectionVersionOrder, where *ent.TemplateSectionVersionWhereInput, themeTemplateID uint64) (*ent.TemplateSectionVersionConnection, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.TemplateSectionVersionUseCase.ListByThemeTemplateID(gc, themeTemplateID, after, first, before, last, orderBy, where)
}
