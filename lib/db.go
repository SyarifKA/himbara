package lib

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "user:password@tcp(127.0.0.1:3306)/himbara?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Gagal konek ke database: %v", err))
	}

	fmt.Println("Berhasil konek ke database MySQL: himbara")
}

func GetDB() *gorm.DB {
	return DB
}
