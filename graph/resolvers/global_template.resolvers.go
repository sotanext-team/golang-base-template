package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app-api/ent"
	"app-api/graph/generated"
	"context"
	"fmt"
)

func (r *globalTemplateResolver) Shop(ctx context.Context, obj *ent.GlobalTemplate) (*ent.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

// GlobalTemplate returns generated.GlobalTemplateResolver implementation.
func (r *Resolver) GlobalTemplate() generated.GlobalTemplateResolver {
	return &globalTemplateResolver{r}
}

type globalTemplateResolver struct{ *Resolver }
