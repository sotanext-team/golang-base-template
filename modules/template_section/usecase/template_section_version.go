package usecase

import (
	"context"

	"app-api/db"
	"app-api/ent"
	"app-api/models"
	"app-api/modules/template_section/repository"
	"app-api/modules/template_section/types"
)

type templateSectionVersionUseCase struct {
	templateSectionVersionRepo repository.TemplateSectionVersionRepository
}

func NewTemplateSectionVersionUseCase() TemplateSectionVersionUseCase {
	templateSectionVersionRepo := repository.NewTemplateSectionVersionRepository()
	return &templateSectionVersionUseCase{
		templateSectionVersionRepo: templateSectionVersionRepo,
	}
}

type TemplateSectionVersionUseCase interface {
	ListByThemeTemplateID(
		ctx context.Context,
		themeTemplateID uint64,
		after *ent.Cursor,
		first *int,
		before *ent.Cursor,
		last *int,
		orderBy *ent.TemplateSectionVersionOrder,
		where *ent.TemplateSectionVersionWhereInput,
	) (*ent.TemplateSectionVersionConnection, error)
}

func (instance *templateSectionVersionUseCase) ListByThemeTemplateID(
	ctx context.Context,
	themeTemplateID uint64,
	after *ent.Cursor,
	first *int,
	before *ent.Cursor,
	last *int,
	orderBy *ent.TemplateSectionVersionOrder,
	where *ent.TemplateSectionVersionWhereInput,
) (*ent.TemplateSectionVersionConnection, error) {
	params := types.TemplateSectionVersionGraphListInput{
		GraphPagination: models.GraphPagination{
			After:  after,
			First:  first,
			Before: before,
			Last:   last,
		},
		OrderBy: orderBy,
		Where:   where,
	}
	return instance.templateSectionVersionRepo.ListByThemeTemplateIDWithCursorPagination(db.GetClient(), ctx, themeTemplateID, params)
}
