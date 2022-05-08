package custom_component

import (
	"context"

	"app-api/ent"
	"app-api/ent/componentproperty"
	"app-api/ent/customcomponent"
	"app-api/models"
)

type Repository interface {
	Create(client *ent.Client, ctx context.Context, data ent.CreateCustomComponentInput) (*ent.CustomComponent, error)
	CreateProperty(client *ent.Client, ctx context.Context, data ent.CreateComponentPropertyInput) (*ent.ComponentProperty, error)
	Delete(client *ent.Client, ctx context.Context, id uint64) error
	DeleteProperty(client *ent.Client, ctx context.Context, id uint64) error
	Find(*ent.Client, context.Context, QueryParams) ([]*ent.CustomComponent, error)
	GetByID(client *ent.Client, ctx context.Context, shopID *uint64, id uint64, joinProps bool) (*ent.CustomComponent, error)
	GetPropByID(client *ent.Client, ctx context.Context, id uint64, joinComponent bool) (*ent.ComponentProperty, error)
	ListByShopID(client *ent.Client, ctx context.Context, shopID *uint64, params ListParams) (*ent.CustomComponentConnection, error)
	Update(client *ent.Client, ctx context.Context, component *ent.CustomComponent, data ent.UpdateCustomComponentInput) (*ent.CustomComponent, error)
	UpdateProperty(client *ent.Client, ctx context.Context, property *ent.ComponentProperty, data ent.UpdateComponentPropertyInput) (*ent.ComponentProperty, error)
}

type QueryParams struct {
	ShopID    *uint64
	CreatedBy uint64
	Handle    string
	IsDraft   *bool
	Name      string
	JoinProps bool
}

type ListParams struct {
	models.GraphPagination
	OrderBy *ent.CustomComponentOrder
	Where   *ent.CustomComponentWhereInput
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (instance *repository) Create(client *ent.Client, ctx context.Context, data ent.CreateCustomComponentInput) (*ent.CustomComponent, error) {
	return client.CustomComponent.Create().SetInput(data).Save(ctx)
}

func (instance *repository) CreateProperty(client *ent.Client, ctx context.Context, data ent.CreateComponentPropertyInput) (*ent.ComponentProperty, error) {
	return client.ComponentProperty.Create().SetInput(data).Save(ctx)
}

func (instance *repository) Delete(client *ent.Client, ctx context.Context, id uint64) error {
	return client.CustomComponent.DeleteOneID(id).Exec(ctx)
}

func (instance *repository) DeleteProperty(client *ent.Client, ctx context.Context, id uint64) error {
	return client.ComponentProperty.DeleteOneID(id).Exec(ctx)
}

func (instance *repository) Find(client *ent.Client, ctx context.Context, params QueryParams) ([]*ent.CustomComponent, error) {
	query := client.CustomComponent.Query()
	if params.ShopID == nil {
		query = query.Where(customcomponent.ShopIDIsNil())
	} else {
		query = query.Where(customcomponent.ShopID(*params.ShopID))
	}
	if params.CreatedBy > 0 {
		query = query.Where(customcomponent.CreatedBy(params.CreatedBy))
	}
	if params.Name != "" {
		query = query.Where(customcomponent.Name(params.Name))
	}
	if params.Handle != "" {
		query = query.Where(customcomponent.Handle(params.Handle))
	}
	if params.IsDraft != nil {
		query = query.Where(customcomponent.IsDraft(*params.IsDraft))
	}
	if params.JoinProps {
		query = query.WithProps()
	}
	return query.All(ctx)
}

func (instance *repository) GetByID(client *ent.Client, ctx context.Context, shopID *uint64, id uint64, joinProps bool) (*ent.CustomComponent, error) {
	query := client.CustomComponent.Query().Where(
		customcomponent.ID(id),
	)
	if shopID == nil {
		query = query.Where(customcomponent.ShopIDIsNil())
	} else {
		query = query.Where(customcomponent.ShopID(*shopID))
	}
	if joinProps {
		query = query.WithProps()
	}
	return query.First(ctx)
}

func (instance *repository) GetPropByID(client *ent.Client, ctx context.Context, id uint64, joinComponent bool) (*ent.ComponentProperty, error) {
	query := client.ComponentProperty.Query().Where(componentproperty.ID(id))
	if joinComponent {
		query = query.WithComponent()
	}
	return query.Only(ctx)
}

func (instance *repository) ListByShopID(client *ent.Client, ctx context.Context, shopID *uint64, params ListParams) (*ent.CustomComponentConnection, error) {
	query := client.CustomComponent.Query()
	if shopID == nil {
		query = query.Where(customcomponent.ShopIDIsNil())
	} else {
		query = query.Where(customcomponent.ShopID(*shopID))
	}
	var opts []ent.CustomComponentPaginateOption
	if params.Where != nil {
		opts = append(opts, ent.WithCustomComponentFilter(params.Where.Filter))
	}
	if params.OrderBy != nil {
		opts = append(opts, ent.WithCustomComponentOrder(params.OrderBy))
	}
	return query.Paginate(ctx, params.After, params.First, params.Before, params.Last, opts...)
}

func (instance *repository) Update(client *ent.Client, ctx context.Context, component *ent.CustomComponent, data ent.UpdateCustomComponentInput) (*ent.CustomComponent, error) {
	return client.CustomComponent.UpdateOne(component).SetInput(data).Save(ctx)
}

func (instance *repository) UpdateProperty(
	client *ent.Client, ctx context.Context, property *ent.ComponentProperty, data ent.UpdateComponentPropertyInput) (*ent.ComponentProperty, error) {

	return client.ComponentProperty.UpdateOne(property).SetInput(data).Save(ctx)
}
