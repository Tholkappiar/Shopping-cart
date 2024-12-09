package main

import (
	"gin-test/inits"
	"gin-test/models"
	"log"
)

func init() {
	inits.ConnectToDB()
}

func main() {
	err := inits.DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	} else {
		log.Println("Migration successful")
	}
}
