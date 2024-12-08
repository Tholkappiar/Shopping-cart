package main

import (
	"gin-test/controllers"
	"gin-test/inits"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	inits.ConnectToDB()
}


func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
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

	// Order routes
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)
	
    // Cart routes
	r.POST("/carts", controllers.CreateCart)
	r.GET("/carts", controllers.GetCarts)

	r.Run() 
}