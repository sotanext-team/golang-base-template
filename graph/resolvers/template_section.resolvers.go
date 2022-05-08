package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app-api/ent"
	"app-api/graph/generated"
	"app-api/graph/models"
	"app-api/modules/template_section/custom_types/request_input"
	"context"

	esContext "github.com/es-hs/es-helper/context"
)

func (r *mutationResolver) TemplateSectionSave(ctx context.Context, sections []*models.TemplateSectionInput, themeTemplateID uint64, saveType models.SaveType) ([]*ent.TemplateSection, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	params := request_input.TemplateSectionsSaveParams{
		SaveType:        saveType,
		ThemeTemplateID: themeTemplateID,
		Sections:        sections,
	}
	return r.TemplateSectionUseCase.Save(gc, params)
}

func (r *mutationResolver) TemplateSectionRevertToVersion(ctx context.Context, themeTemplateID uint64, versionID uint64) ([]*ent.TemplateSection, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	params := request_input.TemplateSectionsRevertParams{
		ThemeTemplateID: themeTemplateID,
		VersionID:       versionID,
	}
	return r.TemplateSectionUseCase.Revert(gc, params)
}

func (r *queryResolver) TemplateSections(ctx context.Context, themeTemplateID uint64) ([]*ent.TemplateSection, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.TemplateSectionUseCase.ListByThemeTemplateID(gc, themeTemplateID)
}

func (r *queryResolver) TemplateSection(ctx context.Context, id uint64) (*ent.TemplateSection, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.TemplateSectionUseCase.FindByID(gc, id)
}

func (r *templateSectionResolver) Area(ctx context.Context, obj *ent.TemplateSection) (*string, error) {
	return (*string)(&obj.Area), nil
}

// TemplateSection returns generated.TemplateSectionResolver implementation.
func (r *Resolver) TemplateSection() generated.TemplateSectionResolver {
	return &templateSectionResolver{r}
}

type templateSectionResolver struct{ *Resolver }
