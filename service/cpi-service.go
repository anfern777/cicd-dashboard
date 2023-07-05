package service

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CpiService interface {
	FindAll(*gin.Context) ([]entity.CloudProviderIntegration, error)
	Save(*gin.Context, entity.CloudProviderIntegration) error
}

type cpiService struct {
}

func NewCpiService() CpiService {
	return &cpiService{}
}

func (service *cpiService) Save(ctx *gin.Context, cpi entity.CloudProviderIntegration) error {
	session := GetDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
	err := session.Save(&cpi).Error
	return err
}

func (service *cpiService) FindAll(ctx *gin.Context) ([]entity.CloudProviderIntegration, error) {
	var cpis []entity.CloudProviderIntegration
	err := GetDB(ctx).Find(&cpis).Error
	if err != nil {
		return cpis, err
	}
	return cpis, nil
}
