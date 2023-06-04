package service

import (
	"fmt"

	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getDB(c *gin.Context) *gorm.DB {
	return c.Keys["DB"].(*gorm.DB)
}

type CpiService interface {
	FindAll(*gin.Context, ) ([]entity.CloudProviderIntegration, error)
	Save(*gin.Context, entity.CloudProviderIntegration) (entity.CloudProviderIntegration, error)
}

type cpiService struct {

}

func NewCpiService() CpiService {
	return &cpiService{}
}

func (service *cpiService) Save(ctx *gin.Context,  cpi entity.CloudProviderIntegration) (entity.CloudProviderIntegration, error) {
	session := getDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
    err := session.Save(&cpi).Error
    if err != nil {
        return cpi, fmt.Errorf("DB ERROR: CPI could not be saved: %w", err)
    }
    return cpi, nil
}

func (service *cpiService) FindAll(ctx *gin.Context) ([]entity.CloudProviderIntegration, error) {
	var cpis []entity.CloudProviderIntegration
	err := getDB(ctx).Find(&cpis).Error
	if err != nil {
		return cpis, fmt.Errorf("DB ERROR: Could not fetch cpis: %w", err)
	}
	return cpis, nil
}

