package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		databaseUrl = "postgresql://hanif:P@ssw0rd.1@192.168.0.107:5432/go_project?sslmode=disable&TimeZone=Asia/Jakarta"
	}

	dsn := databaseUrl
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect ot database")
	}
	fmt.Println("Connected to database")
}
