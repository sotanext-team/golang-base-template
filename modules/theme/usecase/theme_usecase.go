package usecase

import (
	"context"
	"fmt"
	"time"

	"app-api/db"
	"app-api/ent"
	"app-api/models"
	"app-api/modules/theme/repository"
	"app-api/modules/theme/types"

	"github.com/jinzhu/copier"
)

type themeUseCase struct {
	themeRepo repository.ThemeRepository
	entClient *ent.Client
}

func NewThemeUseCase() ThemeUseCase {
	themeRepo := repository.NewThemePostgres()
	client := db.GetClient()
	return &themeUseCase{
		entClient: client,
		themeRepo: themeRepo,
	}
}

type ThemeUseCase interface {
	ThemeCreate(ctx context.Context, inputTheme ent.CreateThemeInput) (*ent.Theme, error)
	ThemeUpdate(ctx context.Context, id int, inputTheme ent.UpdateThemeInput) (*ent.Theme, error)
	ThemeDelete(ctx context.Context, id int) error
	Duplicate(ctx context.Context, id int) (*ent.Theme, error)
	Listing(
		ctx context.Context,
		after *ent.Cursor,
		first *int,
		before *ent.Cursor,
		last *int,
		orderBy *ent.ThemeOrder,
		where *ent.ThemeWhereInput,
	) (*ent.ThemeConnection, error)
	GetByID(ctx context.Context, id uint64) (*ent.Theme, error)
}

func (instance *themeUseCase) GetByID(ctx context.Context, id uint64) (*ent.Theme, error) {
	var (
		err   error
		theme *ent.Theme
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	theme, err = instance.themeRepo.FindByID(db.GetClient(), ctx, currentShop, id)
	if err != nil {
		return nil, err
	}

	return theme, nil
}

func (instance *themeUseCase) ThemeCreate(ctx context.Context, inputTheme ent.CreateThemeInput) (*ent.Theme, error) {
	var (
		theme ent.Theme
		err   error
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	if err = copier.Copy(&theme, inputTheme); err != nil {
		return &theme, err
	}

	theme.ShopID = currentShop.ID

	tx, _ := db.GetClient().Tx(ctx)
	return instance.themeRepo.Create(tx.Client(), ctx, &theme)
}

func (instance *themeUseCase) ThemeUpdate(ctx context.Context, id int, inputTheme ent.UpdateThemeInput) (*ent.Theme, error) {
	var (
		err                 error
		theme, updatedTheme *ent.Theme
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	theme, err = instance.themeRepo.FindByID(db.GetClient(), ctx, currentShop, uint64(id))
	if err != nil {
		return nil, err
	}

	if err = copier.Copy(&theme, inputTheme); err != nil {
		return nil, err
	}

	if updatedTheme, err = instance.themeRepo.Save(db.GetClient(), ctx, theme); err != nil {
		return nil, err
	}

	return updatedTheme, nil
}
func (instance *themeUseCase) ThemeDelete(ctx context.Context, id int) error {
	var (
		err error
		// theme *ent.Theme
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	_, err = instance.themeRepo.FindByID(db.GetClient(), ctx, currentShop, uint64(id))
	if err != nil {
		return err
	}

	err = instance.themeRepo.Delete(db.GetClient(), ctx, uint64(id))
	if err != nil {
		return err
	}
	return nil
}

func (instance *themeUseCase) Duplicate(ctx context.Context, id int) (*ent.Theme, error) {
	var (
		err             error
		theme, newTheme *ent.Theme
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	theme, err = instance.themeRepo.FindByID(db.GetClient(), ctx, currentShop, uint64(id))
	if err != nil {
		return nil, err
	}

	now := time.Now()
	nsec := now.UnixNano()

	newName := fmt.Sprintf("%s - Copy - %d", theme.Name, nsec)
	theme.Name = newName
	theme.ID = 0

	tx, _ := db.GetClient().Tx(ctx)
	newTheme, err = instance.themeRepo.Create(tx.Client(), ctx, theme)
	if err != nil {
		return nil, err
	}
	return newTheme, nil
}

func (instance *themeUseCase) Listing(
	ctx context.Context,
	after *ent.Cursor,
	first *int,
	before *ent.Cursor,
	last *int,
	orderBy *ent.ThemeOrder,
	where *ent.ThemeWhereInput,
) (*ent.ThemeConnection, error) {
	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	params := types.ThemeGraphListInput{
		GraphPagination: models.GraphPagination{
			After:  after,
			First:  first,
			Before: before,
			Last:   last,
		},
		OrderBy: orderBy,
		Where:   where,
	}

	return instance.themeRepo.ListByShopID(db.GetClient(), ctx, currentShop, params)
}
