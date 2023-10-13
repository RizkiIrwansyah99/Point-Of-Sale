package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "success",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
