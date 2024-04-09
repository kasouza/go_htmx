package controllers

import (
	"fmt"
	users "go_htmx/services"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "saske",
	})
}

func HomeClicked(c *gin.Context) {
	html := ""
	users, _ := users.All()

	for _, user := range users {
		html += fmt.Sprintf("<li>%d: %v</li>", user.Id, user.Login)
	}

	c.String(http.StatusOK, html)
}
