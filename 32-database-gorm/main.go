package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}

	fmt.Println("---------------------------------------")
	// create
	db.Create(&Product{
		Name:  "Notebook",
		Price: 2000,
	})

	fmt.Println("---------------------------------------")
	// create batch
	products2 := []Product{
		{Name: "Xbox", Price: 3500.78},
		{Name: "Playstation 5", Price: 5500.78},
		{Name: "iPad 12 pro", Price: 9500.78},
		{Name: "Macbook pro M1 max", Price: 33500.78},
		{Name: "Avell", Price: 12500.78},
	}
	db.Create(&products2)

	fmt.Println("---------------------------------------")
	fmt.Println("select one")
	var product Product
	db.First(&product, 2) // buscar por id
	db.First(&product, "name = ?", "Playstation 5")
	fmt.Println(product)

	fmt.Println("---------------------------------------")
	fmt.Println("select all")
	var products []Product
	db.Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	fmt.Println("---------------------------------------")
	fmt.Println("buscando com limit e offset")
	db.Limit(2).Offset(2).Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	fmt.Println("---------------------------------------")
	fmt.Println("buscando com limit e explicitando where")
	db.Limit(2).Where("price > ? && name LIKE ?", 7000, "%iPad%").Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	fmt.Println("---------------------------------------")
	fmt.Println("update on item")
	var p Product
	db.First(&p)
	p.Name = p.Name + " updated"
	db.Save(&p)

	fmt.Println("---------------------------------------")
	fmt.Println("delete item")
	var p2 Product
	db.First(&p2)
	db.Delete(&p)

}
