package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaskeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "saske/index.tmpl", gin.H{
		"title": "Saske",
	})
}
