package models

import (
	"appWeb/db"
)

type Product struct {
	id          int
	Name        string
	Description string
	Value       float64
	Amount      int
}

func SelectAllProducts() []Product {
	db := db.ConectDataBase()

	selectAllProducts, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, amount int
		var name, description string
		var value float64

		err = selectAllProducts.Scan(&id, &name, &description, &value, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.id = id
		p.Name = name
		p.Description = description
		p.Value = value
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, value float64, amount int) {
	db := db.ConectDataBase()

	saveProducts, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	saveProducts.Exec(name, description, value, amount)
	defer db.Close()
}

func DeleteProduct(idProduct string) {
	db := db.ConectDataBase()

	excludeProduct, err := db.Prepare("Delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	excludeProduct.Exec(idProduct)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConectDataBase()

	productOld, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productUpdate := Product{}

	for productOld.Next() {
		var id, amount int
		var name, description string
		var value float64

		err = productOld.Scan(&id, &name, &description, &value, &amount)
		if err != nil {
			panic(err.Error())
		}
		productUpdate.id = id
		productUpdate.Name = name
		productUpdate.Description = description
		productUpdate.Amount = amount
		productUpdate.Value = value

	}
	defer db.Close()
	return productUpdate
}

func UpdateProduct(id, amount int, name, description string, value float64) {
	db := db.ConectDataBase()

	UpdateProduct, err := db.Prepare("update produtos ste nome=$1, description=$2, preco=$3, quantidade=$4, where id=$5")
	if err != nil {
		panic(err.Error())
	}
	UpdateProduct.Exec(name, description, value, amount, id)
	defer db.Close()
}
