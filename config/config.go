package config

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:Insider2816.@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Konfigurasi GORM dengan level log
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Mendapatkan sql.DB dari gorm.DB untuk mengatur pool koneksi
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}

	// Konfigurasi Database Connection Pool
	sqlDB.SetMaxIdleConns(10)           // Jumlah koneksi idle yang disimpan dalam pool
	sqlDB.SetMaxOpenConns(100)          // Jumlah maksimum koneksi yang dapat dibuka ke database
	sqlDB.SetConnMaxLifetime(time.Hour) // Durasi maksimum sebuah koneksi dapat digunakan (1 jam)

	log.Println("Database connected successfully with pooling configuration!")
}
