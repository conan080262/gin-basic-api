package usercontroller

import "github.com/gin-gonic/gin"

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"API Version": "0.0.1",
		"message":     "user",
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register",
	})
}
