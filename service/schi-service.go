package service

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SchiService interface {
	FindAll(*gin.Context) ([]entity.SourceCodeHostIntegration, error)
	Save(*gin.Context, entity.SourceCodeHostIntegration) error
}

type schiService struct {
}

func NewSchiService() SchiService {
	return &schiService{}
}

func (service *schiService) FindAll(ctx *gin.Context) ([]entity.SourceCodeHostIntegration, error) {
	var schis []entity.SourceCodeHostIntegration
	err := getDB(ctx).Find(&schis).Error
	if err != nil {
		return schis, err
	}
	return schis, nil
}

func (service *schiService) Save(ctx *gin.Context, schi entity.SourceCodeHostIntegration) error {
	session := getDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
	err := session.Save(&schi).Error
	return err
}
