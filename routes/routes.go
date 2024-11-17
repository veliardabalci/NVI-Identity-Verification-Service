package routes

import (
	"github.com/gin-gonic/gin"
	"nvi-identity-verification/business"
	"nvi-identity-verification/controller"
	"nvi-identity-verification/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	nviService := services.NewNviService()
	nviBusiness := business.NewNviBusiness(nviService)
	nviController := controllers.NewNviController(nviBusiness)
	api := r.Group("/api")
	{
		NviCheckGroup(api, nviController)
	}

	return r
}

func NviCheckGroup(router *gin.RouterGroup, controller *controllers.NviController) {
	check := router.Group("/check")
	{
		check.POST("/turkish", controller.CheckTurkishNviHandler)
		check.POST("/foreign", controller.CheckForeignNviHandler)
	}
}
