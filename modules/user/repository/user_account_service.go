package repository

import (
	"context"
	"time"

	graphModels "app-api/graph/models"

	userRPC "github.com/es-hs/erpc/app-account/v1"
	"github.com/sirupsen/logrus"
)

func (u *userImpl) GetUsersByIDs(userList []uint64) (users []*graphModels.User, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	accountService, err := u.GetAccountService()
	if err != nil {
		return nil, err
	}
	userFromService, err := accountService.ListUser(ctx, &userRPC.ListUserRequest{
		Ids: userList,
	})
	logrus.Info("done get user list")
	if err != nil {
		return nil, err
	}
	for k := range userFromService.Users {
		users = append(users, &graphModels.User{
			ID:       userFromService.Users[k].Id,
			UserName: &userFromService.Users[k].Username,
			Email:    &userFromService.Users[k].Email,
		})
	}
	return users, err
	// get from service
	// client := instance.dbEnt
	// return client.User.Query().Where(user.IDIn())
	// // return client.Shop.Query().Where(shop.DefaultDomainEQ(domain)).First(ctx)

	// return users, err
}

func (u *userImpl) GetuserByNameOrEmail(adminEmail string) (user *graphModels.User, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	accountService, err := u.GetAccountService()
	logrus.Info(u.accountConnection)
	if err != nil {
		return nil, err
	}
	userFromService, err := accountService.GetUserByNameOrEmail(ctx, &userRPC.GetUserByNameOrEmailRequest{
		Email: adminEmail,
	})
	logrus.Info("done get user list")
	if err != nil {
		return nil, err
	}
	user = &graphModels.User{
		ID:       userFromService.Id,
		UserName: &userFromService.Username,
		Email:    &userFromService.Email,
	}
	return user, err
}

func (u *userImpl) AddUserShop(userId uint64, shopID uint64) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	accountService, err := u.GetAccountService()
	logrus.Info(u.accountConnection)
	if err != nil {
		return err
	}
	_, err = accountService.AddUserShop(ctx, &userRPC.AddUserShopRequest{
		UserId: userId,
		ShopId: shopID,
	})
	logrus.Info("done get user list")
	if err != nil {
		return err
	}
	return nil
}
func (u *userImpl) RemoveUserShop(userId uint64, shopID uint64) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	accountService, err := u.GetAccountService()
	logrus.Info(u.accountConnection)
	if err != nil {
		return err
	}
	_, err = accountService.RemoveUserShop(ctx, &userRPC.RemoveUserShopRequest{
		UserId: userId,
		ShopId: shopID,
	})
	logrus.Info("done get user list")
	if err != nil {
		return err
	}
	return nil
}
