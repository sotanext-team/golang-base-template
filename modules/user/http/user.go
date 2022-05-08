package http

import (
	"app-api/ent"
	"app-api/modules/user/usecase"
	utils "app-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(dbEnt *ent.Client, accountConnection *grpc.ClientConn) UserHandler {
	userUseCase := usecase.NewUserUseCase(dbEnt, accountConnection)
	return UserHandler{
		userUseCase: userUseCase,
	}
}

func (instance *UserHandler) GetCurrentUser(c *gin.Context) {
	currentUser, currentShop, err := utils.GetAuthenInfo(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"err":    err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"user":   currentUser,
		"shop":   currentShop,
	})
}
