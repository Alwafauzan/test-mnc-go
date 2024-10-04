package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:123@tcp(localhost:3306)/test_mnc_go"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	db.AutoMigrate(&User{})

	DB = db
}