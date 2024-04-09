package main

import (
	"go_htmx/users"
	"go_htmx/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("SECRET_REPLACE_LATER"))
	r.Use(sessions.Sessions("default", store))

	files, err := utils.GetFiles("./templates", ".tmpl")
	if err != nil {
		panic(err)
	}

	r.LoadHTMLFiles(files...)
	r.Static("/assets", "./assets")

	route(r)

	r.Run(":8084")
}

func route(r *gin.Engine) {
	usersGroup := r.Group("users")
	users.Route(usersGroup)
}
