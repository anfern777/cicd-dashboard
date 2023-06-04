package controller

import (
	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
)

type SchiController interface {
	FindAll(*gin.Context) []entity.SourceCodeHostIntegration
	Save(*gin.Context) entity.SourceCodeHostIntegration
}

type schiController struct {
	service service.SchiService
}

func NewSchiController(service service.SchiService) SchiController {
	return &schiController{
		service: service,
	}
}

func (controller *schiController) FindAll(ctx *gin.Context) []entity.SourceCodeHostIntegration {
	schis, _ := controller.service.FindAll(ctx)
	return schis
}

func (controller *schiController) Save(ctx *gin.Context) entity.SourceCodeHostIntegration {
	var user entity.SourceCodeHostIntegration
	ctx.BindJSON(&user)
	controller.service.Save(ctx, user)
	return user
}