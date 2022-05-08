package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app-api/ent"
	"app-api/graph/generated"
	"context"
	"fmt"
)

func (r *entityResolver) FindCustomComponentByID(ctx context.Context, id uint64) (*ent.CustomComponent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindShopByID(ctx context.Context, id uint64) (*ent.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
