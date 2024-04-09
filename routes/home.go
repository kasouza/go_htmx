package routes

import (
	"go_htmx/controllers"

	"github.com/gin-gonic/gin"
)

func Home(r *gin.Engine) {
	home := r.Group("/")
	home.GET("/", controllers.HomeIndex)
	home.POST("clicked", controllers.HomeClicked)
}
