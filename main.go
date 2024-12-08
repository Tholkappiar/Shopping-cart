package main

import (
	"gin-test/controllers"
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

	r.POST("/users", controllers.CreateUser)
	r.POST("/users/login", controllers.LoginUser)
	r.GET("/users", controllers.GetUsers)

	// Item routes
	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.GetItems)

	// Order routes
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)
	
    // Cart routes
	r.POST("/carts", controllers.CreateCart)
	r.GET("/carts", controllers.GetCarts)

	r.Run() 
}