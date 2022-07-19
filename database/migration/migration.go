package migration

import (
	"fmt"
	"go-fiber-api/database"
	"go-fiber-api/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}