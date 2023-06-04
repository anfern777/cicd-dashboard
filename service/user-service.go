package service

import (
	"fmt"

	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	Save(*gin.Context, entity.User) (entity.User, error)
	FindAll(*gin.Context) ([]entity.User, error)
	FindByEmail(*gin.Context, string) (entity.User, error)
}

type userService struct {
}

func NewUserService() (UserService) {
	return &userService{}
}

func (service *userService) Save(ctx *gin.Context, user entity.User) (entity.User, error) {
	session := getDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
    err := session.Save(&user).Error
    if err != nil {
        return user, fmt.Errorf("DB ERROR: User could not be saved: %w", err)
    }
    return user, nil
}

func (service *userService) FindAll(ctx *gin.Context) ([]entity.User, error) {
	var users []entity.User
	err := getDB(ctx).Find(&users).Error
	if err != nil {
		return users, fmt.Errorf("DB ERROR: Could not fetch users: %w", err)
	}
	return users, nil
}

func (service *userService) FindByEmail(ctx *gin.Context, email string) (entity.User, error) {
	var user entity.User
	err := getDB(ctx).First(&user, email).Error
	if err != nil {
		return user, fmt.Errorf("DB ERROR: Could not fetch user: %w", err)
	}
	return user, nil
}