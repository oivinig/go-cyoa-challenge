package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oivinig/go-cyoa-challenge/story"
)

func DisplayIntro(c *gin.Context) {
	data := story.LoadData()
	c.HTML(http.StatusOK, "index.tmpl", data["intro"])
}

func Action(c *gin.Context) {
	data := story.LoadData()
	route := c.Params.ByName("route")
	c.HTML(http.StatusOK, "index.tmpl", data[route])
}
