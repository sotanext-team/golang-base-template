package repository

import (
	"time"

	"app-api/configs"
	"app-api/ent"
	graphModels "app-api/graph/models"

	"github.com/es-hs/erpc"
	userRPC "github.com/es-hs/erpc/app-account/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type userImpl struct {
	dbEnt             *ent.Client
	accountConnection *grpc.ClientConn
	accountService    userRPC.UserClient
}

func NewUserRepository(dbEnt *ent.Client, accountConnection *grpc.ClientConn) UserRepository {
	return &userImpl{
		dbEnt:             dbEnt,
		accountConnection: accountConnection,
	}
}

type UserRepository interface {
	GetUsersByIDs(userList []uint64) (users []*graphModels.User, err error)
	GetuserByNameOrEmail(adminEmail string) (user *graphModels.User, err error)
	AddUserShop(userId uint64, shopID uint64) (err error)
	RemoveUserShop(userId uint64, shopID uint64) (err error)
}

func (u *userImpl) GetAccountService() (userRPC.UserClient, error) {
	if u.accountService == nil {
		var err error
		u.accountConnection, err = erpc.GetConnection(configs.GRPC.Server.Account, 2*time.Second)
		if err != nil {
			// TODO: Sentry
			logrus.Error(err)
		}
		u.accountService = userRPC.NewUserClient(u.accountConnection)
	}
	return u.accountService, nil
}
