package routes

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func Run() {
	getRoutes()
	router.Run(":5000")
}

func getRoutes() {
	//Group 1 images
	v1 := router.Group("/v1")
	addImagesRoute(v1)
}
