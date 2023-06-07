package controller

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
)

type SchiController interface {
	FindAll(*gin.Context) ([]entity.SourceCodeHostIntegration, error)
	Save(*gin.Context) error
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

func (controller *schiController) Save(ctx *gin.Context) error {
	var user entity.SourceCodeHostIntegration
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	err = controller.service.Save(ctx, user)
	return err
}
