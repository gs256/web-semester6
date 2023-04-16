package database

import (
	"fmt"
	"lab5/orders"
	"lab5/products"
	"lab5/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDsn() string {
	host := "0.0.0.0"
	port := "5432"
	user := "postgres"
	pass := "postgres"
	dbname := "postgres"
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
}

func Migrate() {
	// _ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(GetDsn()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// if (!db.Migrator().HasTable(&Product{})) {
	// 	db.Migrator().CreateTable(&Product{})
	// }

	// Migrate the schema
	err = db.AutoMigrate(&products.ProductModel{}, &orders.OrderModel{}, &users.UserModel{}, &orders.OrderProductsModel{})
	if err != nil {
		panic("cant migrate db")
	}
}
