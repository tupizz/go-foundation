package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	GormProducts []GormProduct
}

// GormProduct M <- 1 Category
type GormProduct struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber `gorm:"foreignKey:ProductID;references:id"`
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func findAllProductsJoinCategory(db *gorm.DB) {
	var products []GormProduct
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}

func initProducts(db *gorm.DB) {
	// create batch
	categories := []Category{
		{Name: "Games"},
		{Name: "Apple"},
		{Name: "Notebook"},
	}
	db.Create(&categories)

	// create batch
	products := []GormProduct{
		{Name: "Xbox", Price: 3500.78, CategoryID: categories[0].ID},
		{Name: "Playstation 5", Price: 5500.78, CategoryID: categories[0].ID},
		{Name: "iPad 12 pro", Price: 9500.78, CategoryID: categories[1].ID},
		{Name: "Macbook pro M1 max", Price: 33500.78, CategoryID: categories[1].ID},
		{Name: "Avell", Price: 12500.78, CategoryID: categories[2].ID},
	}
	db.Create(&products)

	// create serial number
	serialNumbers := []SerialNumber{
		{Number: "ASHBAS23", ProductID: products[0].ID},
		{Number: "JFHBAS23", ProductID: products[1].ID},
		{Number: "SDDFRF23", ProductID: products[2].ID},
		{Number: "WECSEA23", ProductID: products[3].ID},
	}
	db.Create(&serialNumbers)
}

func findAllCategories(db *gorm.DB) {
	var categories []Category
	db.Model(&Category{}).Preload("GormProducts.SerialNumber").Find(&categories)
	for _, category := range categories {
		fmt.Printf("Category: %v\n", category.Name)
		for _, product := range category.GormProducts {
			fmt.Printf("Product: %s, Serial Number: %s\n", product.Name, product.SerialNumber.Number)

		}
		fmt.Println("---------------------------------------")
	}
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&GormProduct{}, &Category{}, &SerialNumber{})
	if err != nil {
		return
	}

	fmt.Println("---------------------------------------")

	//initProducts(db)
	//findAllProductsJoinCategory(db)
	findAllCategories(db)
}
