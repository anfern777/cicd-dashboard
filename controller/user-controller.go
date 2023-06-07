package controller

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/anfern777/cicd-dashboard/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	FindAll(*gin.Context) ([]entity.User, error)
	FindByEmail(*gin.Context, string) (entity.User, error)
	Save(ctx *gin.Context) error
}

type userController struct {
	service service.UserService
}

var validate *validator.Validate

func NewUserController(service service.UserService) UserController {
	validate = validator.New()
	validate.RegisterValidation("is-role", validators.ValidateUserRole)
	return &userController{
		service: service,
	}
}

func (controller *userController) FindAll(ctx *gin.Context) ([]entity.User, error) {
	users, err := controller.service.FindAll(ctx)
	return users, err
}

func (controller *userController) FindByEmail(ctx *gin.Context, email string) (entity.User, error) {
	user, err := controller.service.FindByEmail(ctx, email)
	return user, err
}

func (controller *userController) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	err = validate.Struct(user)
	if err != nil {
		return err
	}

	err = controller.service.Save(ctx, user)
	return err
}
