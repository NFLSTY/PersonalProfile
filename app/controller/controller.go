package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	// Dependent services
}

func NewController() *Controller {
	return &Controller{}
}

func (r *Controller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
