package inits

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=ep-silent-night-a1qkabd7.ap-southeast-1.aws.neon.tech user=neondb_owner password=sJx1aiMpG4Nc dbname=shopping-cart port=5432 sslmode=require"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil  {
		log.Fatal("Failed to connect to the Database")
	}
}