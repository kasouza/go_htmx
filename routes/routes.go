package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	Home(r)
	Saske(r)
}
