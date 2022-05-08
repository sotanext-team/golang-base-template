package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app-api/ent"
	"app-api/graph/generated"
	"app-api/graph/models"
	"context"
	"fmt"

	esContext "github.com/es-hs/es-helper/context"
)

func (r *mutationResolver) CreateShop(ctx context.Context, input ent.CreateShopInput) (*ent.Shop, error) {
	return ent.FromContext(ctx).Shop.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateShop(ctx context.Context, id uint64, input ent.UpdateShopInput) (*ent.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateShops(ctx context.Context, ids []uint64, input ent.UpdateShopInput) ([]*ent.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ShopAddManager(ctx context.Context, adminEmail string) (*bool, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	result, err := r.ShopUseCase.AddCurrentShopManager(gc, adminEmail)
	return &result, err
}

func (r *mutationResolver) ShopRemoveManager(ctx context.Context, adminEmail string) (*bool, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	result, err := r.ShopUseCase.RemoveCurrentShopManager(gc, adminEmail)
	return &result, err
}

func (r *queryResolver) Shops(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ShopOrder, where *ent.ShopWhereInput) (*ent.ShopConnection, error) {
	return r.Client.Shop.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithShopOrder(orderBy),
			ent.WithShopFilter(where.Filter),
		)
}

func (r *queryResolver) ShopListManager(ctx context.Context) ([]*models.User, error) {
	gc, err := esContext.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.ShopUseCase.ListCurrentShopAdmin(gc)
	// panic(fmt.Errorf("not implemented"))
}

func (r *shopResolver) Themes(ctx context.Context, obj *ent.Shop) ([]*ent.Theme, error) {
	panic(fmt.Errorf("not implemented"))
}

// Shop returns generated.ShopResolver implementation.
func (r *Resolver) Shop() generated.ShopResolver { return &shopResolver{r} }

type shopResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) ShopAddManager(ctx context.Context, adminEmail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) ShopRemoveManager(ctx context.Context, adminEmail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) ShopListManager(ctx context.Context, adminEmail string) ([]*models.User, error) {
	return r.ShopUseCase.ListCurrentShopAdmin(ctx)
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) ShopRemoveAdmin(ctx context.Context, adminEmail string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}
