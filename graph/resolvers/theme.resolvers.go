package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app-api/ent"
	"app-api/graph/generated"
	"context"
	"fmt"

	esContext "github.com/es-hs/es-helper/context"
)

func (r *mutationResolver) ThemeInsertFromGlobal(ctx context.Context, globalThemeID int) (*ent.Theme, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ThemeDuplicate(ctx context.Context, id int) (*ent.Theme, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeUseCase.Duplicate(gc, id)
}

func (r *mutationResolver) ThemeUpdate(ctx context.Context, id int, theme ent.UpdateThemeInput) (*ent.Theme, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeUseCase.ThemeUpdate(gc, id, theme)
}

func (r *mutationResolver) ThemeDelete(ctx context.Context, id int) (*string, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	err = r.ThemeUseCase.ThemeDelete(gc, id)
	if err != nil {
		return nil, err
	}
	msg := "success"
	return &msg, nil
}

func (r *mutationResolver) ThemeCreate(ctx context.Context, theme ent.CreateThemeInput) (*ent.Theme, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeUseCase.ThemeCreate(gc, theme)
}

func (r *queryResolver) Themes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ThemeOrder, where *ent.ThemeWhereInput) (*ent.ThemeConnection, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ThemeUseCase.Listing(gc, after, first, before, last, orderBy, where)
}

func (r *themeResolver) Shop(ctx context.Context, obj *ent.Theme) (*ent.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *themeResolver) ThemeTemplates(ctx context.Context, obj *ent.Theme) ([]*ent.ThemeTemplate, error) {
	panic(fmt.Errorf("not implemented"))
}

// Theme returns generated.ThemeResolver implementation.
func (r *Resolver) Theme() generated.ThemeResolver { return &themeResolver{r} }

type themeResolver struct{ *Resolver }
