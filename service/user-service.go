package service

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	Save(*gin.Context, entity.User) error
	FindAll(*gin.Context) ([]entity.User, error)
	FindByEmail(*gin.Context, string) (entity.User, error)

	GetAllUserRoles() []string
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Save(ctx *gin.Context, user entity.User) error {
	session := getDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
	err := session.Save(&user).Error
	return err
}

func (service *userService) FindAll(ctx *gin.Context) ([]entity.User, error) {
	var users []entity.User
	err := getDB(ctx).Find(&users).Error
	return users, err
}

func (service *userService) FindByEmail(ctx *gin.Context, email string) (entity.User, error) {
	var user entity.User
	err := getDB(ctx).First(&user, email).Error
	return user, err
}

func (service *userService) GetAllUserRoles() []string {
	return []string{"guest", "admin"}
}
