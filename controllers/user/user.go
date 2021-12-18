package usercontroller

import (
	"net/http"

	"github.com/conan080262/gin-basic-api.git/configs"
	"github.com/conan080262/gin-basic-api.git/models"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"API Version": "0.0.1",
		"message":     "user",
	})
}

func Register(c *gin.Context) {
	var inputJson InputRegister
	if err := c.ShouldBindJSON(&inputJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: inputJson.Fullname,
		Email:    inputJson.Email,
		Password: inputJson.Password,
	}

	result := configs.DB.Debug().Create(&user)

	if result.Error != nil {
		if result.RowsAffected == 0 {

		} else if result.RowsAffected == 1 {

		}
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(201, gin.H{
		"data": user,
		"row":  result.RowsAffected,
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
