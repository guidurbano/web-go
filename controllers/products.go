package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web-go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvertedToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error in conversion of price", err)
		}

		quantityConvertedToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error in conversion of quantity", err)
		}

		models.CreateNewProduct(name, description,
			priceConvertedToFloat, quantityConvertedToInt)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.EditProduct(id)
	temp.ExecuteTemplate(w, "Edit", product)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvertedToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error in the conversion of ID to int:", err)
		}

		priceConvertedToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error in conversion of price", err)
		}

		quantityConvertedToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error in conversion of quantity", err)
		}
		models.UpdateProduct(idConvertedToInt, name, description,
			priceConvertedToFloat, quantityConvertedToInt)
	}

	http.Redirect(w, r, "/", 301)

}
