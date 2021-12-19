package main

import (
	"os"

	"github.com/conan080262/gin-basic-api.git/configs"
	v1 "github.com/conan080262/gin-basic-api.git/routes/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//p4.2:41:18
	router := SetupRouter()
	router.Run(":" + os.Getenv("GO_PORT")) // listen and serve on 0.0.0.0:8080 http://127.0.0.1:3001/ http://localhost:3001
	//gin -a 3002 -p 3000
}

func SetupRouter() *gin.Engine {
	//Load .env
	godotenv.Load(".env")

	configs.Connection()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
		// AllowOrigins:     []string{"https://foo.com"},
		// AllowMethods:     []string{"PUT", "PATCH"},
		// AllowHeaders:     []string{"Origin"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))

	apiV1 := router.Group("/api/v1")
	v1.InitHomeRoutes(apiV1)
	v1.InitUserRoutes(apiV1)
	return router
}
