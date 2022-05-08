package custom_component

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"app-api/configs"
	"app-api/constants"
	"app-api/ent"
	graphmodels "app-api/graph/models"
	"app-api/models"
	"github.com/es-hs/erpc"
	pb "github.com/es-hs/erpc/component"
	"github.com/es-hs/es-helper/errors"
	eslog "github.com/es-hs/es-helper/log"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UseCase interface {
	BuildPropsStr(ctx context.Context, component *ent.CustomComponent) string
	CreateComponent(client *ent.Client, ctx context.Context, shopID *uint64, createdBy uint64, name string, content, entryFileName *string) (*ent.CustomComponent, error)
	CreateProperty(client *ent.Client, ctx context.Context, shopID *uint64, data ent.CreateComponentPropertyInput) (*ent.ComponentProperty, error)
	DeleteProperty(client *ent.Client, ctx context.Context, shopID *uint64, id uint64) error
	GetByID(client *ent.Client, ctx context.Context, shopID *uint64, id uint64, joinProps bool) (*ent.CustomComponent, error)
	GetPropByID(client *ent.Client, ctx context.Context, id uint64, joinComponent bool) (*ent.ComponentProperty, error)
	List(client *ent.Client, ctx context.Context, shopID *uint64, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CustomComponentOrder, where *ent.CustomComponentWhereInput) (*ent.CustomComponentConnection, error)
	PrepareDevComponent(client *ent.Client, ctx context.Context, shopID *uint64, userID uint64) (*ent.CustomComponent, error)
	PrepareDevSession(client *ent.Client, ctx context.Context, shopID *uint64, userID uint64, componentID uint64) (*graphmodels.DevSession, error)
	ReloadLivePreview(ctx context.Context, userID uint64, component *ent.CustomComponent) error
	UpdateComponent(client *ent.Client, ctx context.Context, component *ent.CustomComponent, data ent.UpdateCustomComponentInput) (*ent.CustomComponent, error)
	UpdateAndReload(client *ent.Client, ctx context.Context, shopID *uint64, id uint64, data *ent.UpdateCustomComponentInput, userID uint64) (*ent.CustomComponent, error)
	UpdateProperty(client *ent.Client, ctx context.Context, shopID *uint64, id uint64, data ent.UpdateComponentPropertyInput) (*ent.ComponentProperty, error)
}

type useCase struct {
	repo Repository
}

func NewUseCase() UseCase {
	return &useCase{
		repo: NewRepository(),
	}
}

// BuildPropsStr returns a string represent component's properties for usages in HTML template
func (instance *useCase) BuildPropsStr(ctx context.Context, component *ent.CustomComponent) string {
	var (
		props       []*ent.ComponentProperty
		builtParams []string
		err         error
		result      string
	)
	props, err = component.Edges.PropsOrErr()
	if ent.IsNotLoaded(err) {
		props = component.QueryProps().AllX(ctx)
	}
	for _, prop := range props {
		builtParams = append(builtParams, fmt.Sprintf(`%s="%s"`, prop.Name, prop.Value))
	}
	result = strings.Join(builtParams, " ")
	return result
}

// CreateComponent creates a new custom component if a component with the same name does not exist, otherwise it will return
// the existing component
func (instance *useCase) CreateComponent(
	client *ent.Client, ctx context.Context, shopID *uint64, createdBy uint64, name string, content, entryFileName *string) (*ent.CustomComponent, error) {

	if err := validateComponentName(name); err != nil {
		return nil, err
	}
	if entryFileName == nil || *entryFileName == "" {
		n := fmt.Sprintf("%s.js", name)
		entryFileName = &n
	}
	if err := validateEntryFileName(*entryFileName); err != nil {
		return nil, err
	}

	results, err := instance.repo.Find(client, ctx, QueryParams{
		ShopID: shopID,
		Name:   name,
	})
	if err != nil {
		return nil, err
	}
	if len(results) > 0 {
		return results[0], nil
	}

	isDraft := false
	return instance.repo.Create(client, ctx, ent.CreateCustomComponentInput{
		ShopID:        shopID,
		CreatedBy:     createdBy,
		Name:          name,
		Content:       content,
		EntryFileName: entryFileName,
		Handle:        newComponentHandle(),
		IsDraft:       &isDraft,
	})
}

func (instance *useCase) CreateProperty(client *ent.Client, ctx context.Context, shopID *uint64, data ent.CreateComponentPropertyInput) (*ent.ComponentProperty, error) {
	component, err := instance.GetByID(client, ctx, shopID, data.ComponentID, false)
	if err != nil {
		return nil, err
	}
	if component == nil {
		return nil, errors.NewNotExistsError("component does not exist")
	}
	if data.Name == "" {
		return nil, errors.NewValidationError("name is required")
	}

	// TODO: handle live reloading
	return instance.repo.CreateProperty(client, ctx, data)
}

func (instance *useCase) DeleteProperty(client *ent.Client, ctx context.Context, shopID *uint64, id uint64) error {
	prop, err := instance.GetPropByID(client, ctx, id, true)
	if err != nil {
		return err
	}
	if prop == nil {
		return nil
	}
	if prop.Edges.Component.ShopID != shopID {
		return errors.NewValidationError("you don't have access to this component")
	}

	// TODO: handle live reloading
	return instance.repo.DeleteProperty(client, ctx, id)
}

// GetByID gets custom component of a shop using its ID. It will return nil if no component found
func (instance *useCase) GetByID(
	client *ent.Client, ctx context.Context, shopID *uint64, id uint64, joinProps bool) (*ent.CustomComponent, error) {

	result, err := instance.repo.GetByID(client, ctx, shopID, id, joinProps)
	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	} else {
		return result, err
	}
}

func (instance *useCase) GetPropByID(client *ent.Client, ctx context.Context, id uint64, joinComponent bool) (*ent.ComponentProperty, error) {
	result, err := instance.repo.GetPropByID(client, ctx, id, joinComponent)
	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	} else {
		return result, err
	}
}

func (instance *useCase) List(
	client *ent.Client, ctx context.Context, shopID *uint64, after *ent.Cursor, first *int, before *ent.Cursor, last *int,
	orderBy *ent.CustomComponentOrder, where *ent.CustomComponentWhereInput) (*ent.CustomComponentConnection, error) {

	return instance.repo.ListByShopID(client, ctx, shopID, ListParams{
		GraphPagination: models.GraphPagination{
			After:  after,
			First:  first,
			Before: before,
			Last:   last,
		},
		OrderBy: orderBy,
		Where:   where,
	})
}

func (instance *useCase) PrepareDevComponent(client *ent.Client, ctx context.Context, shopID *uint64, userID uint64) (*ent.CustomComponent, error) {
	component, err := instance.getDraftComponent(client, ctx, shopID, userID)
	if err != nil {
		return nil, fmt.Errorf("get draft component: %w", err)
	}
	if component == nil {
		// If user has no draft component, create a new default component
		component, err = instance.newDefaultComponent(client, ctx, shopID, userID)
		if err != nil {
			return nil, fmt.Errorf("new default component: %w", err)
		}
	}
	return component, nil
}

func (instance *useCase) PrepareDevSession(client *ent.Client, ctx context.Context, shopID *uint64, userID uint64, componentID uint64) (*graphmodels.DevSession, error) {
	var (
		component *ent.CustomComponent
		err       error
	)
	if componentID > 0 {
		component, err = instance.GetByID(client, ctx, shopID, componentID, true)
		if err != nil {
			return nil, fmt.Errorf("get component: %w", err)
		}
		if component == nil {
			return nil, errors.NewNotExistsError("component does not exist")
		}
		// TODO: handle cases that session has been terminated, or session belongs to another user
	} else {
		// If request has no component ID, prepare a new component for this dev session
		component, err = instance.PrepareDevComponent(client, ctx, shopID, userID)
		if err != nil {
			return nil, fmt.Errorf("prepare component: %w", err)
		}
	}

	// Prepare to request component service for a dev session
	conn, err := erpc.GetConnection(configs.GRPC.Server.Component, configs.GRPC.ConnectTimeout)
	if err != nil {
		return nil, fmt.Errorf("grpc conn: %w", err)
	}
	sessionClient := pb.NewSessionClient(conn)

	rpcCtx, cancelFunc := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFunc()
	sid := uint64(0)
	if shopID != nil {
		sid = *shopID
	}
	c, err := ToGRPCObject(component)
	if err != nil {
		return nil, err
	}
	// Request a dev session from component service
	result, err := sessionClient.PrepareDevSession(rpcCtx, &pb.PrepareDevSessionRequest{
		ShopId:    sid,
		UserId:    userID,
		Component: c,
	})
	if err != nil {
		return nil, fmt.Errorf("prepare dev session: %w", err)
	}
	// Update component with new session ID
	sessID := result.GetSession().GetId()
	component, err = instance.UpdateComponent(client, ctx, component, ent.UpdateCustomComponentInput{
		SessionID: &sessID,
	})

	return &graphmodels.DevSession{
		ID:         result.GetSession().GetId(),
		PreviewURL: result.GetSession().GetPreviewUrl(),
		Component:  component,
	}, nil
}

func (instance *useCase) ReloadLivePreview(ctx context.Context, userID uint64, component *ent.CustomComponent) error {
	conn, err := erpc.GetConnection(configs.GRPC.Server.Component, configs.GRPC.ConnectTimeout)
	if err != nil {
		return err
	}
	c, err := ToGRPCObject(component)
	if err != nil {
		return err
	}
	compClient := pb.NewComponentClient(conn)
	_, err = compClient.UpdateLivePreview(ctx, &pb.UpdateLivePreviewRequest{
		BuildDist: false,
		UserId:    userID,
		Component: c,
	})
	return err
}

func (instance *useCase) UpdateComponent(
	client *ent.Client, ctx context.Context, component *ent.CustomComponent, data ent.UpdateCustomComponentInput) (*ent.CustomComponent, error) {

	return instance.repo.Update(client, ctx, component, data)
}

func (instance *useCase) UpdateAndReload(
	client *ent.Client, ctx context.Context, shopID *uint64, id uint64, data *ent.UpdateCustomComponentInput, userID uint64) (*ent.CustomComponent, error) {

	component, err := instance.GetByID(client, ctx, shopID, id, true)
	if err != nil {
		return nil, err
	}
	if component == nil {
		return nil, errors.NewNotExistsError("component does not exist")
	}
	if data != nil {
		component, err = instance.UpdateComponent(client, ctx, component, *data)
		if err != nil {
			return nil, err
		}
	}

	// TODO: check if current session belongs to this user or not

	go func() {
		// Request component service to update live preview for this component
		ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancelFunc()
		err := instance.ReloadLivePreview(ctx, userID, component)
		if err != nil {
			errStatus, _ := status.FromError(err)
			if errStatus.Code() == codes.FailedPrecondition && errStatus.Message() == "no active session" {
				// If session was terminated due to inactivity, restart the session
				// TODO: restart inactive session
			} else {
				eslog.LogError(ctx, err)
			}
		}
	}()

	return component, nil
}

func (instance *useCase) UpdateProperty(client *ent.Client, ctx context.Context, shopID *uint64, id uint64, data ent.UpdateComponentPropertyInput) (*ent.ComponentProperty, error) {
	prop, err := instance.GetPropByID(client, ctx, id, true)
	if err != nil {
		return nil, err
	}
	if prop == nil {
		return nil, errors.NewNotExistsError("property does not exist")
	}
	if prop.Edges.Component.ShopID != shopID {
		return nil, errors.NewValidationError("you don't have access to this component")
	}

	// TODO: handle live reloading
	return instance.repo.UpdateProperty(client, ctx, prop, data)
}

// getDraftComponent returns a component with IsDraft = true that was created for a specific user
func (instance *useCase) getDraftComponent(
	client *ent.Client, ctx context.Context, shopID *uint64, createdBy uint64) (*ent.CustomComponent, error) {

	isDraft := true
	results, err := instance.repo.Find(client, ctx, QueryParams{
		ShopID:    shopID,
		CreatedBy: createdBy,
		IsDraft:   &isDraft,
	})
	if err != nil {
		return nil, err
	} else if len(results) > 0 {
		return results[0], nil
	} else {
		return nil, nil
	}
}

// newDefaultComponent creates a component with default attributes in the DB. Usually for starting dev session.
func (instance *useCase) newDefaultComponent(
	client *ent.Client, ctx context.Context, shopID *uint64, createdBy uint64) (*ent.CustomComponent, error) {

	isDraft := true
	efn := constants.DefaultEntryFileName
	c := constants.NewComponentContent
	return instance.repo.Create(client, ctx, ent.CreateCustomComponentInput{
		ShopID:        shopID,
		CreatedBy:     createdBy,
		Name:          constants.NewComponentName,
		EntryFileName: &efn,
		Handle:        newComponentHandle(),
		Content:       &c,
		IsDraft:       &isDraft,
	})
}

func newComponentHandle() string {
	// Use UUID to create component handle
	return uuid.New().String()
}

func validateComponentName(name string) error {
	reg, _ := regexp.Compile("^[a-zA-Z][a-zA-Z0-9]+$")
	if !reg.MatchString(name) {
		return errors.NewValidationError(
			"component name can only contain alpha numeric characters, and can not start with a number",
		)
	}
	return nil
}

func validateEntryFileName(name string) error {
	reg, _ := regexp.Compile("^[a-zA-Z][a-zA-Z0-9]+.js$")
	if !reg.MatchString(name) {
		return errors.NewValidationError(
			"entry file name can only contain alpha numeric characters, and can not start with a number. E.g: remoteEntry.js",
		)
	}
	return nil
}
