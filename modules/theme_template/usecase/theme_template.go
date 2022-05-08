package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"app-api/db"
	"app-api/ent"
	"app-api/ent/themetemplate"
	"app-api/models"
	_templateSectionModuleRepo "app-api/modules/template_section/repository"
	"app-api/modules/theme_template/repository"
	"app-api/modules/theme_template/types"

	grpcLibs "app-api/libs/grpc"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type themeTemplateUseCase struct {
	themeTemplateRepo   repository.ThemeTemplateRepository
	globalTemplateRepo  repository.GlobalTemplateRepository
	templateSectionRepo _templateSectionModuleRepo.TemplateSectionRepository
}

func NewThemeTemplateUseCase() ThemeTemplateUseCase {
	themeTemplateRepo := repository.NewThemeTemplateRepository()
	globalTemplateRepo := repository.NewGlobalTemplateRepository()
	templateSectionRepo := _templateSectionModuleRepo.NewTemplateSectionRepository()
	return &themeTemplateUseCase{
		themeTemplateRepo:   themeTemplateRepo,
		globalTemplateRepo:  globalTemplateRepo,
		templateSectionRepo: templateSectionRepo,
	}
}

type ThemeTemplateUseCase interface {
	Create(ctx context.Context, inputData ent.CreateThemeTemplateInput) (*ent.ThemeTemplate, error)
	Update(ctx context.Context, id uint64, inputData ent.UpdateThemeTemplateInput) (*ent.ThemeTemplate, error)
	ListByThemeTemplateID(
		ctx context.Context,
		themeId uint64,
		pageType *themetemplate.PageType,
		after *ent.Cursor,
		first *int,
		before *ent.Cursor,
		last *int,
		orderBy *ent.ThemeTemplateOrder,
		where *ent.ThemeTemplateWhereInput,
	) (*ent.ThemeTemplateConnection, error)
	ListByThemeTemplateIDInTrash(
		ctx context.Context,
		themeId uint64,
		pageType *themetemplate.PageType,
		after *ent.Cursor,
		first *int,
		before *ent.Cursor,
		last *int,
		orderBy *ent.ThemeTemplateOrder,
		where *ent.ThemeTemplateWhereInput,
	) (*ent.ThemeTemplateConnection, error)
	Delete(ctx context.Context, id uint64) error
	ForceDelete(ctx context.Context, id uint64) error
	Duplicate(ctx context.Context, id uint64) (*ent.ThemeTemplate, error)
	MakeDefault(ctx context.Context, id uint64) (*ent.ThemeTemplate, error)
	Restore(ctx context.Context, id uint64) (*ent.ThemeTemplate, error)
	GetByID(ctx context.Context, id uint64) (*ent.ThemeTemplate, error)
	MakeGlobal(ctx context.Context, themeTemplateId uint64) (*ent.GlobalTemplate, error)
	Publish(ctx context.Context, themeTemplateId uint64) (*ent.ThemeTemplate, error)
	// InsertFromGlobal(ctx context.Context, globalTemplateId uint64) (*ent.ThemeTemplate, error)
}

func (instance *themeTemplateUseCase) Create(ctx context.Context, inputData ent.CreateThemeTemplateInput) (*ent.ThemeTemplate, error) {
	var (
		err              error
		themeTemplate    ent.ThemeTemplate
		newThemeTemplate *ent.ThemeTemplate
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	_, err = instance.themeTemplateRepo.FindByName(db.GetClient(), ctx, currentShop, inputData.Name, *inputData.ThemeID, *inputData.PageType)

	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if !ent.IsNotFound(err) {
		return nil, errors.New("name is exist")
	}

	if err = copier.Copy(&themeTemplate, inputData); err != nil {
		return nil, err
	}

	themeTemplate.ShopID = currentShop.ID

	if newThemeTemplate, err = instance.themeTemplateRepo.Create(db.GetClient(), ctx, &themeTemplate); err != nil {
		return nil, err
	}

	return newThemeTemplate, nil
}

func (instance *themeTemplateUseCase) Update(ctx context.Context, id uint64, inputData ent.UpdateThemeTemplateInput) (*ent.ThemeTemplate, error) {
	var (
		err                  error
		themeTemplate        ent.ThemeTemplate
		updatedThemeTemplate *ent.ThemeTemplate
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	_, err = instance.themeTemplateRepo.FindByID(db.GetClient(), ctx, currentShop, id)
	if err != nil {
		return nil, err
	}

	// Find exist name
	if *inputData.Name != themeTemplate.Name {
		_, err := instance.themeTemplateRepo.FindByName(db.GetClient(), ctx, currentShop, *inputData.Name, *inputData.ThemeID, *inputData.PageType)
		if err != nil && !ent.IsNotFound(err) {
			logrus.Info("hello 1")
			return nil, err
		}
		if !ent.IsNotFound(err) {
			return nil, errors.New("name is exist")
		}
	}

	if err = copier.Copy(&themeTemplate, inputData); err != nil {
		logrus.Info("hello 2")
		return nil, err
	}

	themeTemplate.ID = id
	if updatedThemeTemplate, err = instance.themeTemplateRepo.Save(db.GetClient(), ctx, &themeTemplate); err != nil {
		return nil, err
	}

	return updatedThemeTemplate, nil
}

func (instance *themeTemplateUseCase) ListByThemeTemplateID(
	ctx context.Context,
	themeId uint64,
	pageType *themetemplate.PageType,
	after *ent.Cursor,
	first *int,
	before *ent.Cursor,
	last *int,
	orderBy *ent.ThemeTemplateOrder,
	where *ent.ThemeTemplateWhereInput,
) (*ent.ThemeTemplateConnection, error) {
	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	params := types.ThemeTemplateGraphListInput{
		GraphPagination: models.GraphPagination{
			After:  after,
			First:  first,
			Before: before,
			Last:   last,
		},
		OrderBy: orderBy,
		Where:   where,
	}
	return instance.themeTemplateRepo.ListByThemeIDAndPageTypeWithCursorPagination(db.GetClient(), ctx, currentShop, uint(themeId), *pageType, params)
}

func (instance *themeTemplateUseCase) ListByThemeTemplateIDInTrash(
	ctx context.Context,
	themeId uint64,
	pageType *themetemplate.PageType,
	after *ent.Cursor,
	first *int,
	before *ent.Cursor,
	last *int,
	orderBy *ent.ThemeTemplateOrder,
	where *ent.ThemeTemplateWhereInput,
) (*ent.ThemeTemplateConnection, error) {
	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	params := types.ThemeTemplateGraphListInput{
		GraphPagination: models.GraphPagination{
			After:  after,
			First:  first,
			Before: before,
			Last:   last,
		},
		OrderBy: orderBy,
		Where:   where,
	}
	return instance.themeTemplateRepo.ListByThemeIDAndPageTypeWithCursorPaginationIntrash(db.GetClient(), ctx, currentShop, uint(themeId), *pageType, params)
}

func (instance *themeTemplateUseCase) Delete(ctx context.Context, id uint64) error {
	var (
		err           error
		themeTemplate *ent.ThemeTemplate
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	themeTemplate, err = instance.themeTemplateRepo.FindByID(db.GetClient(), ctx, currentShop, id)
	if err != nil {
		return err
	}

	if !ent.IsNotFound(err) {
		err = instance.themeTemplateRepo.Delete(db.GetClient(), ctx, themeTemplate.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (instance *themeTemplateUseCase) ForceDelete(ctx context.Context, id uint64) error {
	var (
		err           error
		themeTemplate *ent.ThemeTemplate
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	themeTemplate, err = instance.themeTemplateRepo.FindByIDUnscoped(db.GetClient(), ctx, currentShop, id)
	if err != nil {
		return err
	}

	if themeTemplate.ID == 0 {
		return errors.New("record not found")
	}

	err = instance.themeTemplateRepo.ForceDelete(db.GetClient(), ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (instance *themeTemplateUseCase) Duplicate(ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	var (
		err                             error
		themeTemplate, newThemeTemplate *ent.ThemeTemplate
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	themeTemplate, err = instance.themeTemplateRepo.FindByID(db.GetClient(), ctx, currentShop, id)
	if err != nil {
		return nil, err
	}

	if themeTemplate.ID == 0 {
		return nil, errors.New("record not found")
	}

	now := time.Now()
	nsec := now.UnixNano()

	newName := fmt.Sprintf("%s - Copy - %d", themeTemplate.Name, nsec)
	themeTemplate.Name = newName
	themeTemplate.ID = 0

	newThemeTemplate, err = instance.themeTemplateRepo.Create(db.GetClient(), ctx, themeTemplate)
	if err != nil {
		return nil, err
	}
	return newThemeTemplate, nil
}

func (instance *themeTemplateUseCase) MakeDefault(ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	var (
		err                                 error
		themeTemplate, updatedThemeTemplate *ent.ThemeTemplate
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	themeTemplate, err = instance.themeTemplateRepo.FindByID(db.GetClient(), ctx, currentShop, id)
	if err != nil {
		return nil, err
	}

	if themeTemplate.ID == 0 {
		return nil, errors.New("record not found")
	}

	err = db.WithTx(ctx, db.GetClient(), func(tx *ent.Tx) error {
		client := tx.Client()
		// Update field default of all others to false
		err = instance.themeTemplateRepo.SetDefault(client, ctx, currentShop, themeTemplate.ThemeID, false)
		if err != nil {
			return err
		}

		// Update the current is true
		themeTemplate.Default = true
		updatedThemeTemplate, err = instance.themeTemplateRepo.Save(client, ctx, themeTemplate)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedThemeTemplate, nil
}

func (instance *themeTemplateUseCase) Restore(ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	var (
		err                                 error
		themeTemplate, updatedThemeTemplate *ent.ThemeTemplate
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	themeTemplate, err = instance.themeTemplateRepo.FindByIDUnscoped(db.GetClient(), ctx, currentShop, id)
	if err != nil {
		return nil, err
	}

	if themeTemplate.ID == 0 {
		return nil, errors.New("record not found")
	}

	if updatedThemeTemplate, err = instance.themeTemplateRepo.Restore(db.GetClient(), ctx, id); err != nil {
		return nil, err
	}

	return updatedThemeTemplate, nil
}

func (instance *themeTemplateUseCase) GetByID(ctx context.Context, id uint64) (*ent.ThemeTemplate, error) {
	var (
		err           error
		themeTemplate *ent.ThemeTemplate
	)

	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	themeTemplate, err = instance.themeTemplateRepo.FindByID(db.GetClient(), ctx, currentShop, id)
	if err != nil {
		return nil, err
	}

	return themeTemplate, nil
}

func (instance *themeTemplateUseCase) MakeGlobal(ctx context.Context, id uint64) (*ent.GlobalTemplate, error) {
	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	var (
		globalTemplate, newGlobalTemplate *ent.GlobalTemplate
		err                               error
	)

	_, err = instance.globalTemplateRepo.FindByThemeTemplateId(db.GetClient(), ctx, currentShop, id)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if !ent.IsNotFound(err) {
		return nil, errors.New("this template already made global")
	}

	themeTemplate, err := instance.themeTemplateRepo.FindByID(db.GetClient(), ctx, currentShop, id)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	nsec := now.UnixNano()

	newName := fmt.Sprintf("%s - Global - %d", themeTemplate.Name, nsec)

	globalTemplate = &ent.GlobalTemplate{
		ShopID: currentShop.ID,
		// ThemeTemplateID: uint(themeTemplateId),
		Name:         newName,
		ViewCount:    0,
		InstallCount: 0,
	}

	newGlobalTemplate, err = instance.globalTemplateRepo.Create(db.GetClient(), ctx, globalTemplate)
	if err != nil {
		return nil, err
	}

	return newGlobalTemplate, nil
}

func (instance *themeTemplateUseCase) Publish(ctx context.Context, themeTemplateId uint64) (*ent.ThemeTemplate, error) {
	shopValue := ctx.Value("currentShop")
	currentShop := shopValue.(*ent.Shop)

	// Call DeployShop
	err := grpcLibs.DeployServiceDeployShop(currentShop.DefaultDomain, "ap-southeast-1", "", "", "")
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// func (instance *themeTemplateUseCase) InsertFromGlobal(ctx context.Context, globalTemplateId uint64) (*ent.ThemeTemplate, error) {
// 	shopValue := ctx.Value("currentShop")
// 	currentShop := shopValue.(*ent.Shop)

// 	var (
// 		existGlobalTemplate *ent.GlobalTemplate
// 		err                 error
// 	)

// 	existGlobalTemplate, err = instance.globalTemplateRepo.FindByThemeTemplateId(instance.database, currentShop, uint(globalTemplateId))
// 	if err != nil {
// 		return models.ThemeTemplate{}, err
// 	}

// 	themeTemplate, err := instance.themeTemplateRepo.FindByID(instance.database, currentShop, existGlobalTemplate.ThemeTemplateID)
// 	if err != nil {
// 		return models.ThemeTemplate{}, err
// 	}

// 	templateSections, err := instance.templateSectionRepo.ListByThemeTemplateID(instance.database, int(themeTemplate.ID))
// 	if err != nil {
// 		return models.ThemeTemplate{}, err
// 	}

// 	// Duplicate themeTemplate and templateSections
// 	err = instance.database.Transaction(func(tx *gorm.DB) error {
// 		themeTemplate.ID = 0

// 		for index := range templateSections {
// 			templateSections[index].ID = 0

// 			themeTemplate.TemplateSections = append(themeTemplate.TemplateSections, &templateSections[index])
// 		}

// 		if err = instance.themeTemplateRepo.Create(tx, &themeTemplate); err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	return themeTemplate, nil
// }
