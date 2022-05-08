package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app-api/ent"
	"app-api/ent/themetemplate"
	"app-api/graph/generated"
	"context"
	"fmt"

	esContext "github.com/es-hs/es-helper/context"
)

func (r *mutationResolver) ThemeTemplateCreate(ctx context.Context, themeTemplate ent.CreateThemeTemplateInput) (*ent.ThemeTemplate, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeTemplateUseCase.Create(gc, themeTemplate)
}

func (r *mutationResolver) ThemeTemplateUpdate(ctx context.Context, id uint64, themeTemplate ent.UpdateThemeTemplateInput) (*ent.ThemeTemplate, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.ThemeTemplateUseCase.Update(gc, id, themeTemplate)
}

func (r *mutationResolver) ThemeTemplateDelete(ctx context.Context, id uint64) (*string, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	err = r.ThemeTemplateUseCase.Delete(gc, id)
	if err != nil {
		return nil, err
	}
	msg := "success"
	return &msg, nil
}

func (r *mutationResolver) ThemeTemplateDuplicate(ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeTemplateUseCase.Duplicate(gc, id)
}

func (r *mutationResolver) ThemeTemplateMakeDefault(ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeTemplateUseCase.MakeDefault(gc, id)
}

func (r *mutationResolver) ThemeTemplateRestore(ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeTemplateUseCase.Restore(gc, id)
}

func (r *mutationResolver) ThemeTemplateForceDelete(ctx context.Context, id uint64) (*string, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	err = r.ThemeTemplateUseCase.ForceDelete(gc, id)
	if err != nil {
		return nil, err
	}
	msg := "success"
	return &msg, nil
}

func (r *mutationResolver) ThemeTemplateMakeGlobal(ctx context.Context, id uint64) (*ent.GlobalTemplate, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeTemplateUseCase.MakeGlobal(gc, id)
}

func (r *mutationResolver) ThemeTemplateInsertFromGlobal(ctx context.Context, globalTemplateID uint64) (*ent.ThemeTemplate, error) {
	// gc, err := esContext.GinContextFromContext(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ThemeTemplatePublish(ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeTemplateUseCase.Publish(gc, id)
}

func (r *queryResolver) ThemeTemplates(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ThemeTemplateOrder, where *ent.ThemeTemplateWhereInput, pageType *themetemplate.PageType, themeID uint64) (*ent.ThemeTemplateConnection, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeTemplateUseCase.ListByThemeTemplateID(gc, themeID, pageType, after, first, before, last, orderBy, where)
}

func (r *queryResolver) ThemeTemplatesTrash(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ThemeTemplateOrder, where *ent.ThemeTemplateWhereInput, pageType *themetemplate.PageType, themeID uint64) (*ent.ThemeTemplateConnection, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeTemplateUseCase.ListByThemeTemplateIDInTrash(gc, themeID, pageType, after, first, before, last, orderBy, where)
}

func (r *queryResolver) ThemeTemplate(ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeTemplateUseCase.GetByID(gc, id)
}

func (r *themeTemplateResolver) Theme(ctx context.Context, obj *ent.ThemeTemplate) (*ent.Theme, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeUseCase.GetByID(gc, obj.ThemeID)
}

func (r *themeTemplateResolver) GlobalTemplate(ctx context.Context, obj *ent.ThemeTemplate) (*ent.GlobalTemplate, error) {
	panic(fmt.Errorf("not implemented"))
}

// ThemeTemplate returns generated.ThemeTemplateResolver implementation.
func (r *Resolver) ThemeTemplate() generated.ThemeTemplateResolver { return &themeTemplateResolver{r} }

type themeTemplateResolver struct{ *Resolver }
