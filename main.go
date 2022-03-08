package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/callback", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
