package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API Version": "0.0.1",
			"message":     "pong",
		})
	})
	r.Run(":3001") // listen and serve on 0.0.0.0:8080 http://127.0.0.1:3001/ http://localhost:3001
	//gin -a 3001 -p 3000
}
