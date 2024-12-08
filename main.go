package main

import (
	"gin-test/test1"
	"gin-test/test2"

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

	r.POST("/test1", test1.Test1)
	r.POST("/test2", test2.Test2)
	// r.GET("/users", controllers.GetUsers)

	// // Item routes
	// r.POST("/items", controllers.CreateItem)
	// r.GET("/items", controllers.GetItems)

	// // Cart routes
	// r.POST("/carts", controllers.CreateCart)
	// r.GET("/carts", controllers.GetCarts)

	// // Order routes
	// r.POST("/orders", controllers.CreateOrder)
	// r.GET("/orders", controllers.GetOrders)

	r.Run() 
}