package routes_web

import (
	"github.com/gin-gonic/gin"
	"my-go/app/controllers"
)

func Init(router *gin.Engine )  {
	router.GET("/", controllers.Index)
}