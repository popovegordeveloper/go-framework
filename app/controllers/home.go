package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"net/http"
)

func Index(c *gin.Context) {
	h := gin.H{}
	h["csrf_token"] = csrf.Token(c.Request)
	c.Status(http.StatusOK)

	Render.Execute("pages/index", h, c.Request, c.Writer)
}
