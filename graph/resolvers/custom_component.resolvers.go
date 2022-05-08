package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app-api/db"
	"app-api/ent"
	"app-api/graph/models"
	"context"
	"fmt"

	eserrors "github.com/es-hs/es-helper/errors"
	eslog "github.com/es-hs/es-helper/log"
)

func (r *mutationResolver) CustomComponentCreate(ctx context.Context, input ent.CreateCustomComponentInput) (*ent.CustomComponent, error) {
	var (
		shopID *uint64
		userID uint64
	)

	// TODO: get shopID and userID from context
	return r.CustomComponentUseCase.CreateComponent(db.GetClient(), ctx, shopID, userID, input.Name, input.Content, input.EntryFileName)
}

func (r *mutationResolver) CustomComponentUpdate(ctx context.Context, id uint64, input ent.UpdateCustomComponentInput) (*ent.CustomComponent, error) {
	var (
		component *ent.CustomComponent
		err       error
		shopID    *uint64
		dbClient  = db.GetClient()
	)
	// TODO: get shopID from context
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}
	if input.Name == nil && input.Content == nil {
		return nil, fmt.Errorf("either name or content must be provided")
	}

	component, err = r.CustomComponentUseCase.GetByID(dbClient, ctx, shopID, id, true)
	if err != nil {
		return nil, fmt.Errorf("get component: %w", err)
	}
	if component == nil {
		return nil, fmt.Errorf("component does not exists")
	}
	component, err = r.CustomComponentUseCase.UpdateComponent(dbClient, ctx, component, input)
	if err != nil {
		eslog.LogError(ctx, err)
		return nil, fmt.Errorf("update component: %w", err)
	}
	return component, nil
}

func (r *mutationResolver) CustomComponentUpdateLive(ctx context.Context, id uint64, input *ent.UpdateCustomComponentInput) (*ent.CustomComponent, error) {
	var (
		component *ent.CustomComponent
		err       error
		shopID    *uint64
		userID    uint64
		valError  *eserrors.NotExistsError
		dbClient  = db.GetClient()
	)
	// TODO: get shopID and userID from context
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}
	if input != nil && input.Name == nil && input.Content == nil {
		return nil, fmt.Errorf("either name or content must be provided")
	}
	component, err = r.CustomComponentUseCase.UpdateAndReload(dbClient, ctx, shopID, id, input, userID)
	if eserrors.As(err, &valError) {
		return nil, err
	} else if err != nil {
		eslog.LogError(ctx, err)
		return nil, fmt.Errorf("update component live: %w", err)
	}
	return component, nil
}

func (r *mutationResolver) PrepareDevSession(ctx context.Context, input models.PrepareDevSessionInput) (*models.DevSession, error) {
	var (
		shopID      *uint64
		userID      uint64
		componentID uint64
		notExist    *eserrors.NotExistsError
		dbClient    = db.GetClient()
	)
	// TODO: get shopID and userID from context

	if input.ComponentID != nil {
		componentID = *input.ComponentID
	}
	devSession, err := r.CustomComponentUseCase.PrepareDevSession(dbClient, ctx, shopID, userID, componentID)
	if eserrors.As(err, &notExist) {
		return nil, err
	} else if err != nil {
		eslog.LogError(ctx, err)
		return nil, err
	}
	return devSession, nil
}

func (r *queryResolver) CustomComponents(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CustomComponentOrder, where *ent.CustomComponentWhereInput) (*ent.CustomComponentConnection, error) {
	var shopID *uint64
	// TODO: get shopID from context

	return r.CustomComponentUseCase.List(db.GetClient(), ctx, shopID, after, first, before, last, orderBy, where)
}
