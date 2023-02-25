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

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		amount := r.FormValue("amount")
		value := r.FormValue("value")

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro de converção Id", err)
		}
		valueConverted, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println("Erro de converção value", err)
		}
		amountConverted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Erro de converção amount", err)
		}

		models.UpdateProduct(idConverted, amountConverted, name, description, valueConverted)
	}
	http.Redirect(w, r, "/", 301)
}
