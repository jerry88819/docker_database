package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	
)

func Init() (db *gorm.DB, err error) {

	// postgres:123@tcp(127.0.0.1:3306)/postgres?charset=utf8mb4&parseTime=True&loc=Local
	// pgUrl := "postgresql://postgres:123@127.0.0.1/postgres?sslmode=require"

	connectStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSL_MODE"))

	// pgUrl := "host=127.0.0.1 user=postgres password=123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Taipei"

	fmt.Println(connectStr)
	
	db, err = gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Println("Open pg db error:", err)
		return
	}

	_db, err := db.DB()
	if err != nil {
		log.Println("Set db connection pool failed:", err)
		return
	}

	_db.SetMaxIdleConns(2)
	_db.SetMaxOpenConns(5)
	_db.SetConnMaxLifetime(time.Hour)

	return
} // Init()
