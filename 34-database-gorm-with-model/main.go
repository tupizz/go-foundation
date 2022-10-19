package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	GormProducts []Product `gorm:"many2many:product_categories;"`
}

// Product M <- N Category
type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
	gorm.Model
}

func initProducts(db *gorm.DB) {
	// create batch
	categories := []Category{
		{Name: "Games"},
		{Name: "Apple"},
		{Name: "Notebook"},
		{Name: "Electronics"},
	}
	db.Create(&categories)

	// create batch
	products := []Product{
		{Name: "Xbox", Price: 3500.78, Categories: []Category{categories[0], categories[3]}},
		{Name: "Playstation 5", Price: 5500.78, Categories: []Category{categories[0], categories[3]}},
		{Name: "iPad 12 pro", Price: 9500.78, Categories: []Category{categories[1], categories[2], categories[3]}},
		{Name: "Macbook pro M1 max", Price: 33500.78, Categories: []Category{categories[1], categories[2], categories[3]}},
		{Name: "Avell", Price: 12500.78, Categories: []Category{categories[1], categories[2], categories[3]}},
	}
	db.Create(&products)

}

func selectAllProductsJoinCategory(db *gorm.DB) {
	var products []Product
	//db.Model(&Product{}).Preload("Categories").Find(&products) // outra forma de fazer
	db.Preload("Categories").Find(&products)
	for _, product := range products {
		fmt.Printf("Product: %v with Categories:\n", product.Name)
		for _, category := range product.Categories {
			fmt.Printf("\tcategory: %v\n", category.Name)
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

	err = db.AutoMigrate(&Product{}, &Category{})
	if err != nil {
		return
	}

	fmt.Println("---------------------------------------")

	//initProducts(db)
	selectAllProductsJoinCategory(db)
}
