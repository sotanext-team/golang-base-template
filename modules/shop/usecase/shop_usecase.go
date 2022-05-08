package usecase

import (
	"context"
	"errors"

	"app-api/ent"
	graphModels "app-api/graph/models"
	"app-api/models"
	"app-api/modules/shop/repository"
	userRepo "app-api/modules/user/repository"

	authz "github.com/es-hs/authzclient"
	"google.golang.org/grpc"

	"app-api/pkg/utils"

	"github.com/sirupsen/logrus"
)

type shopUseCase struct {
	shopRepo    repository.ShopRepository
	userRepo    userRepo.UserRepository
	databaseENT *ent.Client
}

type Shop struct {
	ID            uint64
	DefaultDomain string
	ShopName      string
}

func NewShopUseCase(dbEnt *ent.Client, accountConnection *grpc.ClientConn) ShopUseCase {
	shopRepo := repository.NewShopRepository(dbEnt)
	userRepo := userRepo.NewUserRepository(dbEnt, accountConnection)
	return &shopUseCase{
		shopRepo: shopRepo,
		userRepo: userRepo,
	}
}

type ShopUseCase interface {
	ListCurrentShopAdmin(ctx context.Context) (users []*graphModels.User, err error)
	AddCurrentShopManager(ctx context.Context, adminEmail string) (result bool, err error)
	RemoveCurrentShopManager(ctx context.Context, adminEmail string) (result bool, err error)
	CreateShop(ctx context.Context, inputShop models.Shop) (models.Shop, error)
	CreateShopRevert(ctx context.Context, inputShop models.Shop) error
	GetShop(ctx context.Context, id uint64) (Shop, error)
	GetShopByDomain(ctx context.Context, domain string) (Shop, error)
}

func (instance *shopUseCase) CreateShop(ctx context.Context, inputShop models.Shop) (models.Shop, error) {
	shop, err := instance.shopRepo.CreateShop(ctx, inputShop)
	if err != nil {
		return models.Shop{}, err
	}
	// Call DeployShop
	// err = grpcLibs.DeployServiceDeployShop(shop.DefaultDomain, "ap-southeast-1", "", "", "")
	if err != nil {
		return models.Shop{}, err
	}
	return shop, nil
}

func (instance *shopUseCase) CreateShopRevert(ctx context.Context, inputShop models.Shop) error {
	return instance.shopRepo.CreateShopRevert(ctx, inputShop)
}

func (instance *shopUseCase) ListCurrentShopAdmin(ctx context.Context) (users []*graphModels.User, err error) {
	// userValue := ctx.Value("currentUser")
	// currentUser := userValue.(*graphModels.User)
	// shopValue := ctx.Value("currentShop")
	// currentShop := shopValue.(*ent.Shop)
	currentUser, currentShop, err := utils.GetAuthenInfo(ctx)
	if err != nil {
		return nil, err
	}
	// authz.AddRoleToDomain(uint64(currentUser.ID), uint64(currentShop.ID), authz.LOGIN_PERMISSION)
	// _, _ = authz.GenerateOwnerRole(currentUserID, currentShop.ID)
	// logrus.Info(result1)
	// logrus.Info(err1)
	result, err := authz.CheckPermission(currentUser.ID, currentShop.ID, authz.LOGIN_PERMISSION)
	if !result || err != nil {
		logrus.Info(currentUser.ID)
		logrus.Info(currentShop.ID)
		logrus.Info(result)
		logrus.Info(err)
		return nil, errors.New("You don't have owner permission or something wrong")
	}

	userList, _, err := authz.GetAllUsersByDomain(currentShop.ID)

	if users, err = instance.userRepo.GetUsersByIDs(userList); err != nil {
		return nil, err
	}
	for k := range users {

		roles, _ := authz.GetRolesInDomain(uint64(users[k].ID), uint64(currentShop.ID))
		for i := range roles {
			users[k].Roles = append(users[k].Roles, &roles[i])
		}
		// users = append(users, u)
	}
	return users, nil
}

func (instance *shopUseCase) AddCurrentShopManager(ctx context.Context, adminEmail string) (bool, error) {
	// userValue := ctx.Value("currentUser")
	// currentUser := userValue.(*graphModels.User)
	// shopValue := ctx.Value("currentShop")
	// currentShop := shopValue.(*ent.Shop)
	currentUser, currentShop, err := utils.GetAuthenInfo(ctx)
	if err != nil {
		return false, err
	}
	result, err := authz.CheckPermission(currentUser.ID, currentShop.ID, authz.OWNER_ROLE)
	if !result || err != nil {
		return false, errors.New("You don't have owner permission or something wrong")
	}

	user, err := instance.userRepo.GetuserByNameOrEmail(adminEmail)
	if err != nil {
		return false, err
	}

	// TODO we should u SAGA here
	// add role of admin to shop domain
	_, err = authz.AddRoleToDomain(user.ID, currentShop.ID, authz.ADMIN_ROLE)
	if err != nil {
		return false, err
	}
	// add user_shop to user service
	err = instance.userRepo.AddUserShop(user.ID, currentShop.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (instance *shopUseCase) RemoveCurrentShopManager(ctx context.Context, adminEmail string) (bool, error) {
	// userValue := ctx.Value("currentUser")
	// currentUser := userValue.(*graphModels.User)
	// shopValue := ctx.Value("currentShop")
	// currentShop := shopValue.(*ent.Shop)
	currentUser, currentShop, err := utils.GetAuthenInfo(ctx)
	if err != nil {
		return false, err
	}
	result, err := authz.CheckPermission(currentUser.ID, currentShop.ID, authz.OWNER_ROLE)
	if !result || err != nil {
		return false, errors.New("You don't have owner permission or something wrong")
	}

	user, err := instance.userRepo.GetuserByNameOrEmail(adminEmail)
	if err != nil {
		return false, err
	}

	// TODO we should u SAGA here
	// add role of admin to shop domain
	_, err = authz.RemoveRoleFromDomain(user.ID, currentShop.ID, authz.ADMIN_ROLE)
	if err != nil {
		return false, err
	}
	// add user_shop to user service
	err = instance.userRepo.RemoveUserShop(user.ID, currentShop.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (instance *shopUseCase) GetShop(ctx context.Context, id uint64) (Shop, error) {
	shop, err := instance.shopRepo.GetShop(ctx, id)
	if err != nil {
		return Shop{}, err
	}
	return Shop{
		ID:            uint64(shop.ID),
		DefaultDomain: shop.DefaultDomain,
		ShopName:      shop.ShopName,
	}, nil
}

func (instance *shopUseCase) GetShopByDomain(ctx context.Context, domain string) (Shop, error) {
	shop, err := instance.shopRepo.GetShopByDomainE(ctx, domain)
	if err != nil {
		return Shop{}, err
	}
	return Shop{
		ID:            uint64(shop.ID),
		DefaultDomain: shop.DefaultDomain,
		ShopName:      shop.ShopName,
	}, nil
}
