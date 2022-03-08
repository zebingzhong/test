package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/callback", func(c *gin.Context) {
		buf := make([]byte, 1024)

		n, _ := c.Request.Body.Read(buf)

		fmt.Println(string(buf[0:n]))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
