package controller

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll(*gin.Context) []entity.User
	FindByEmail(*gin.Context, string) entity.User
	Save(ctx *gin.Context) entity.User
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (controller *userController) FindAll(ctx *gin.Context) []entity.User {
	users, _ := controller.service.FindAll(ctx)
	return users
}

func (controller *userController) FindByEmail(ctx *gin.Context, email string) entity.User {
	user, _ := controller.service.FindByEmail(ctx, email)
	return user
}

func (controller *userController) Save(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	controller.service.Save(ctx, user)
	return user
}