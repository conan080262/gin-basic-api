package v1

import "github.com/gin-gonic/gin"

func InitHomeRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API Version": "0.0.1",
			"message":     "home",
		})
	})
}
