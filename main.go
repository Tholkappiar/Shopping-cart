package main

import (
	"gin-test/controllers"
	"gin-test/inits"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load Envs and connect to DB
	inits.ConnectToDB()
}


func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is nice , from thols",
		})
	})

	r.POST("/users", controllers.CreateUser)
	r.POST("/users/login", controllers.LoginUser)
	r.GET("/users", controllers.GetUsers)

	// Item routes
	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.GetItems)

	// Cart routes
	r.POST("/carts", controllers.CreateCart)
	r.GET("/carts", controllers.GetCarts)

	// Order routes
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)

	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	r.Run(":"+port) 
}