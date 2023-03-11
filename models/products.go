package models

import (
	"web-go/db"
)

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func GetProducts() []Product {
	db := db.ConnectDB()

	selectProducts, err := db.Query("SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDB()
	insertProducts, err := db.Prepare(
		"INSERT INTO products(name, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProducts.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProducts, err := db.Prepare(
		"DELETE FROM products WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProducts.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDB()

	storeProduct, err := db.Query(
		"SELECT * FROM products WHERE id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}
	for storeProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = storeProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}

	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDB()
	insertProducts, err := db.Prepare(
		"UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	insertProducts.Exec(name, description, price, quantity, id)
	defer db.Close()
}
