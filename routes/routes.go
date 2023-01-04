package routes

import (
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/oivinig/go-cyoa-challenge/controllers"
)

func Handler() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", controllers.DisplayIntro)
	r.GET("/:route", controllers.Action)
	r.Run(":8080")
}
