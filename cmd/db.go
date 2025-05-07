package main

import (
	"fmt"
	"leagueinform/internal/types"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartDatabase() *gorm.DB {
	pw := os.Getenv("DBPASS")
	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/testdb?parseTime=true", pw)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&types.Account{})

	fmt.Println("Connected and migrated")

	return db

}
