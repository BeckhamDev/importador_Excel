package router

import (
	"importador_Excel/internal/controllers"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {

	r.POST("/saveData", controllers.UploadFileAndSaveData)

}
