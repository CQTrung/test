package boostrap

import (
	"trunne-crud/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "trunne:Trunne61@tcp(127.0.0.1:3306)/ladgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	db.AutoMigrate(&entity.Employee{})
	return db
}
