package main

import (
	"github.com/gin-gonic/gin"
	"my-go/routes/api"
	"my-go/routes/web"
)

func setupRouter() *gin.Engine {

	//Initialize routes
	router := gin.Default()
	routes_web.Init(router)
	routes_api.Init(router)

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
