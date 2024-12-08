package migrate

import (
	"gin-test/inits"
	"gin-test/models"
	"log"
)

func init() {
	inits.ConnectToDB()
}

func Migrate() {
	err := inits.DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	} else {
		log.Println("Migration successful")
	}
}