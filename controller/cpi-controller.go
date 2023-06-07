package controller

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
)

type CpiController interface {
	FindAll(ctx *gin.Context) ([]entity.CloudProviderIntegration, error)
	Save(ctx *gin.Context) error
}

type cpiController struct {
	service service.CpiService
}

func NewCpiController(service service.CpiService) CpiController {
	return &cpiController{
		service: service,
	}
}

func (controller *cpiController) FindAll(ctx *gin.Context) ([]entity.CloudProviderIntegration, error) {
	cpis, err := controller.service.FindAll(ctx)
	return cpis, err
}

func (controller *cpiController) Save(ctx *gin.Context) error {
	var user entity.CloudProviderIntegration
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	err = controller.service.Save(ctx, user)
	return err
}
