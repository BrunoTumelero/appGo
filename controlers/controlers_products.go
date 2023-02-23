package controlers

import (
	"appWeb/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SelectAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		value := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		valueConverted, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println("Erro ao converter preço:", err)
		}

		amountConverted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Erro ao converter a quantidade:", err)
		}

		models.CreateNewProduct(name, description, valueConverted, amountConverted)
	}
	http.Redirect(w, r, "/", 301)
}
