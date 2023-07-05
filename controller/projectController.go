package controller

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
)

type ProjectController interface {
	FindAll(ctx *gin.Context) ([]entity.Project, error)
	FindAllByUser(ctx *gin.Context) ([]entity.Project, error)
	Save(ctx *gin.Context) error
}

type projectController struct {
	service service.ProjectService
}

func NewProjectController(service service.ProjectService) ProjectController {
	return &projectController{
		service: service,
	}
}

func (controller *projectController) FindAll(ctx *gin.Context) ([]entity.Project, error) {
	prjs, err := controller.service.FindAll(ctx)
	return prjs, err
}

func (controller *projectController) FindAllByUser(ctx *gin.Context) ([]entity.Project, error) {
	prjs, err := controller.service.FindAllByUser(ctx)
	return prjs, err
}

func (controller *projectController) Save(ctx *gin.Context) error {
	var prj entity.Project
	err := ctx.ShouldBindJSON(&prj)
	if err != nil {
		return err
	}
	err = controller.service.Save(ctx, prj)
	return err
}
