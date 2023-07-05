package service

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProjectService interface {
	FindAll(*gin.Context) ([]entity.Project, error)
	FindAllByUser(ctx *gin.Context) ([]entity.Project, error)
	FindByID(ctx *gin.Context, projectID uint) (entity.Project, error)
	Save(*gin.Context, entity.Project) error
}

type projectService struct {
}

func NewProjectService() ProjectService {
	return &projectService{}
}

func (project *projectService) Save(ctx *gin.Context, cpi entity.Project) error {
	user := ctx.Keys["jwt-user"].(entity.User)
	cpi.UserID = user.ID
	session := GetDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
	err := session.Save(&cpi).Error
	return err
}

func (project *projectService) FindAll(ctx *gin.Context) ([]entity.Project, error) {
	var prjs []entity.Project
	err := GetDB(ctx).Find(&prjs).Error
	if err != nil {
		return prjs, err
	}
	return prjs, nil
}

func (project *projectService) FindAllByUser(ctx *gin.Context) ([]entity.Project, error) {
	user := ctx.Keys["jwt-user"].(entity.User)
	var prjs []entity.Project
	err := GetDB(ctx).Model(&user).Association("Projects").Find(&prjs)
	if err != nil {
		return nil, err
	}

	return prjs, nil
}

func (project *projectService) FindByID(ctx *gin.Context, projectID uint) (entity.Project, error) {
	var prj entity.Project

	err := GetDB(ctx).Find(&prj, projectID).Error
	if err != nil {
		return prj, err
	}

	return prj, nil
}
