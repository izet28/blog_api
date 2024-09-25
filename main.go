package main

import (
	"blog/config"
	"blog/models"
	"blog/routers"
	"log"
)

func main() {
	// Inisialisasi database
	config.InitDB()

	// Migrasi model ke dalam database
	config.DB.AutoMigrate(
		&models.Category{},
		&models.SubCategory{},
		&models.Blog{},
		&models.Tag{},
	)

	// Setup router
	router := routers.SetupRouter()

	// Menjalankan server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
