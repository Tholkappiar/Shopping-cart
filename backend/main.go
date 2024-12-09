package main

import (
	"gin-test/controllers"
	"gin-test/inits"
	"gin-test/models"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	inits.ConnectToDB()
	err := inits.DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.CartItem{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	} else {
		log.Println("Migration successful")
	}
}

func main() {
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// User routes
	r.POST("/users", controllers.CreateUser)
	r.POST("/users/login", controllers.LoginUser)
	r.GET("/users", controllers.GetAllUsers)

	// Item routes
	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.GetItems)

	// Cart routes
	r.POST("/cart", controllers.AddToCart)
	r.GET("/cart", controllers.GetUserCart) // New API to get user-specific cart

	// Order routes
	r.POST("/checkout", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders) // New API to get user-specific orders

	r.Run(":8080") // Start server on port 8080
}
