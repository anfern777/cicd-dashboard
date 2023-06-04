package service

import (
	"fmt"

	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SchiService interface {
	FindAll(*gin.Context) ([]entity.SourceCodeHostIntegration, error)
	Save(*gin.Context, entity.SourceCodeHostIntegration) (entity.SourceCodeHostIntegration, error)
}

type schiService struct {
}

func NewSchiService() SchiService{
	return &schiService{}
}

func (service *schiService) FindAll(ctx *gin.Context) ([]entity.SourceCodeHostIntegration, error) {
	var schis []entity.SourceCodeHostIntegration
	err := getDB(ctx).Find(&schis).Error
	if err != nil {
        return schis, fmt.Errorf("DB ERROR: User could not be saved: %w", err)
    }
    return schis, nil
}

func (service *schiService) Save(ctx *gin.Context, schi entity.SourceCodeHostIntegration) (entity.SourceCodeHostIntegration, error) {
	session := getDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
    err := session.Save(&schi).Error
    if err != nil {
        return schi, fmt.Errorf("DB ERROR: SCHI could not be saved: %w", err)
    }
    return schi, nil
}