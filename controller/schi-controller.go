package controller

import (
	"strconv"

	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
)

type SchiController interface {
	FindAll(*gin.Context) ([]entity.SourceCodeHostIntegration, error)
	FindByProject(ctx *gin.Context) (entity.SourceCodeHostIntegration, error)
	Save(*gin.Context) error
	ListEnvironments(ctx *gin.Context) ([]service.Environment, error)
}

type schiController struct {
	service service.SchiService
}

func NewSchiController(service service.SchiService) SchiController {
	return &schiController{
		service: service,
	}
}

func (controller *schiController) FindAll(ctx *gin.Context) ([]entity.SourceCodeHostIntegration, error) {
	schis, err := controller.service.FindAll(ctx)
	return schis, err
}

func (controller *schiController) FindByProject(ctx *gin.Context) (entity.SourceCodeHostIntegration, error) {
	projectID, err := strconv.Atoi(ctx.Param("project_id"))
	var schi entity.SourceCodeHostIntegration
	if err != nil {
		return schi, err
	}
	schi, err = controller.service.FindByProject(ctx, uint(projectID))
	return schi, err
}

func (controller *schiController) Save(ctx *gin.Context) error {
	var schi entity.SourceCodeHostIntegration
	projectID, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return err
	}
	schi.ProjectID = uint(projectID)

	// TODO: verify that project exists

	err = ctx.ShouldBindJSON(&schi)
	if err != nil {
		return err
	}

	err = controller.service.Save(ctx, schi)
	return err
}

func (controller *schiController) ListEnvironments(ctx *gin.Context) ([]service.Environment, error) {

	projectID, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		return nil, err
	}

	schi, err := controller.service.FindByProject(ctx, uint(projectID))
	if err != nil {
		return nil, err
	}

	envs, err := controller.service.ListEnvironments(ctx, schi)
	if err != nil {
		return nil, err
	}

	return envs, nil
}
