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
	}

}
