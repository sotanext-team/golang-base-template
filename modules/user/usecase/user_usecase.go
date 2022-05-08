package usecase

import (
	"app-api/ent"
	"app-api/modules/user/repository"

	"google.golang.org/grpc"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(dbEnt *ent.Client, accountConnection *grpc.ClientConn) UserUseCase {
	userRepo := repository.NewUserRepository(dbEnt, accountConnection)
	return &userUseCase{
		userRepo: userRepo,
	}
}

type UserUseCase interface {
}
