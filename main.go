package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome  string
	Desc  string
	Preco float64
	Qtd   int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Desc: "Camiseta Azul", Preco: 64.35, Qtd: 5},
		{"Tênis", "Adidas", 100.50, 10},
		{"Fone", "Xiaomi", 59, 2},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}
