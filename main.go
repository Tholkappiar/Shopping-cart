package main

import (
	"gin-test/test1"

	"github.com/gin-gonic/gin"
)

func init() {
}


func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is nice , from thols",
		})
	})

	r.GET("/test1", test1.Test1)

	r.Run() 
}