package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app-api/db"
	"app-api/ent"
	"app-api/graph/generated"
	"context"

	eserrors "github.com/es-hs/es-helper/errors"
	eslog "github.com/es-hs/es-helper/log"
)

func (r *mutationResolver) ComponentPropertyCreate(ctx context.Context, input ent.CreateComponentPropertyInput) (*ent.ComponentProperty, error) {
	var (
		shopID   *uint64
		valError *eserrors.ValidationError
		notExist *eserrors.NotExistsError
	)
	p, err := r.CustomComponentUseCase.CreateProperty(db.GetClient(), ctx, shopID, input)
	if eserrors.As(err, &valError) || eserrors.As(err, &notExist) {
		return nil, err
	} else if err != nil {
		eslog.LogError(ctx, err)
		return nil, err
	}
	return p, nil
}

func (r *mutationResolver) ComponentPropertyUpdate(ctx context.Context, id uint64, input ent.UpdateComponentPropertyInput) (*ent.ComponentProperty, error) {
	var (
		shopID   *uint64
		valError *eserrors.ValidationError
		notExist *eserrors.NotExistsError
	)
	p, err := r.CustomComponentUseCase.UpdateProperty(db.GetClient(), ctx, shopID, id, input)
	if eserrors.As(err, &valError) || eserrors.As(err, &notExist) {
		return nil, err
	} else if err != nil {
		eslog.LogError(ctx, err)
		return nil, err
	}
	return p, nil
}

func (r *mutationResolver) ComponentPropertyRemove(ctx context.Context, id uint64) (*bool, error) {
	var (
		shopID   *uint64
		valError *eserrors.ValidationError
	)
	err := r.CustomComponentUseCase.DeleteProperty(db.GetClient(), ctx, shopID, id)
	ok := err == nil
	if !ok && !eserrors.As(err, &valError) {
		eslog.LogError(ctx, err)
	}
	return &ok, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
