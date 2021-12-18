package v1

import (
	usercontroller "github.com/conan080262/gin-basic-api.git/controllers/user"
	"github.com/conan080262/gin-basic-api.git/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) { //
	// routerGroup := rg.Group("/users").Use(middlewares.AuthJWT())
	routerGroup := rg.Group("/users")

	// users(routerGroup.Group("/test"))
	// {{domain_url}}/api/v1/users
	routerGroup.GET("/", usercontroller.GetAll)
	// {{domain_url}}/api/v1/users/register
	routerGroup.POST("/register", usercontroller.Register)
	// {{domain_url}}/api/v1/users/login
	routerGroup.POST("/login", usercontroller.Login)
	// {{domain_url}}/api/v1/users/10 //Param
	routerGroup.GET("/:id", usercontroller.GetById)
	// {{domain_url}}/api/v1/users/search?fullname=jon&id=2 //Param
	routerGroup.GET("/search", usercontroller.SearchByFullname)

	routerGroup.GET("/me", middlewares.AuthJWT(), usercontroller.GetProfile)
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
