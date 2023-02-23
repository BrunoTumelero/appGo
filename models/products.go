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

	selectAllProducts, err := db.Query("select * from produtos")
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
