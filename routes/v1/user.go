package v1

import (
	usercontroller "github.com/conan080262/gin-basic-api.git/controllers/user"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	// users(routerGroup.Group("/test"))

	// {{domain_url}}/api/v1/users
	routerGroup.GET("/", usercontroller.GetAll)

	// {{domain_url}}/api/v1/users/register
	routerGroup.POST("/register", usercontroller.Register)

}

// func users(rg *gin.RouterGroup) {
// 	routerGroup := rg.Group("/")
// 	routerGroup.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"API Version": "0.0.1",
// 			"message":     "a",
// 		})
// 	})
// }
