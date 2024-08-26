package main

import (
	"PersonalProfile/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.Static("/public", "./public")

	routes.SetupRouter(r)

	r.Run("0.0.0.0:8080")
}
