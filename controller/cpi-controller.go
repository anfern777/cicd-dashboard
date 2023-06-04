package controller

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
)

type CpiController interface {
	FindAll(ctx *gin.Context) []entity.CloudProviderIntegration
	Save(ctx *gin.Context) entity.CloudProviderIntegration
}

type cpiController struct {
	service service.CpiService
}

func NewCpiController(service service.CpiService) CpiController {
	return &cpiController{
		service: service,
	}
}

func (controller *cpiController) FindAll(ctx *gin.Context) []entity.CloudProviderIntegration {
	cpis, _ := controller.service.FindAll(ctx)
	return cpis
}

func (controller *cpiController) Save(ctx *gin.Context) entity.CloudProviderIntegration {
	var user entity.CloudProviderIntegration
	ctx.BindJSON(&user)
	controller.service.Save(ctx, user)
	return user
}