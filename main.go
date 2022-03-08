package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/callback", func(c *gin.Context) {
		json := make(map[string]interface{}) //注意该结构接受的内容
		err := c.BindJSON(&json)
		if err != nil {
			return 
		}
		fmt.Println(json)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
