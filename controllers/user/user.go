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
		"data": "register",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "login",
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"data": id,
	})
}

func SearchByFullname(c *gin.Context) {
	// id := c.DefaultQuery("firstname", "Guest")
	fullname := c.Query("fullname")
	c.JSON(200, gin.H{
		"data": fullname,
	})
}
