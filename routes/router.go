package routes

import (
	"PersonalProfile/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	mainController := controller.NewController()

	r.GET("/", mainController.Index)
	r.GET("/detail", mainController.Detail)
}
