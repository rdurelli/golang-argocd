package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/durelli", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Rafael",
		})
	})

	r.GET("/rafael", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Durelli",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
