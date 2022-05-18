package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addImagesRoute(rg *gin.RouterGroup) {
	images := rg.Group("/images")
	images.GET("/get", getImages)
}
func getImages(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "It Works!",
		"data":    "",
	})
	return
}
