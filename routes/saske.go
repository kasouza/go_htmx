package routes

import (
	"go_htmx/controllers"

	"github.com/gin-gonic/gin"
)

func Saske(r *gin.Engine) {
	saske := r.Group("/saske")
	saske.GET("/", controllers.SaskeIndex)
}
