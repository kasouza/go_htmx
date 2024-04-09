package users

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func loginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "users/login.html.tmpl", gin.H{})
}

func loginPost(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.String(500, "Erro interno do servidor")
		return
	}

	email := c.Request.Form.Get("email")
	pass := c.Request.Form.Get("password")

	user, err := maybeLoginUser(email, pass)
	if err != nil {
		fmt.Println(err.Error())
		c.String(500, "Erro interno do servidor")
		return
	}

	if user == nil {
		c.String(400, "Email ou senha inválidos")
		return
	}

	session := sessions.Default(c)
	session.Set("user", user)

	c.Header("HX-Redirect", "/")
	c.String(200, "Sucesso!")
}

func Route(g *gin.RouterGroup) {
	g.GET("login", loginGet)
	g.POST("login", loginPost)
}
