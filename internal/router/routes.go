package router

import (
	"importador_Excel/internal/controllers"
	"importador_Excel/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {

	//Rotas disponiveis da API
	r.POST("/createUser", controllers.CreateUser)
	r.POST("/login", controllers.Login)

	app := r.Group("/app")
	app.Use(middleware.AuthRequired())
	{
		app.POST("/saveData", controllers.UploadFileAndSaveData)
		app.GET("/customersByName/:customerName", controllers.GetRowByCustomerName)
		app.GET("/customersByMeterCategory/:meterCategory", controllers.GetRowByMeterCategory)
		app.GET("/customersByMeterRegion/:meterRegion", controllers.GetRowByMeterRegion)
		app.GET("/customersByResourceGroup/:resourceGroup", controllers.GetRowByResourceGroup)
	}

}
