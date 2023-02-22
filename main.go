package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func conectDataBase() *sql.DB {
	conn := "user=postgres dbname=loja_go password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	id          int
	Name        string
	Description string
	Value       float64
	Amount      int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8090", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectDataBase()

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

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
