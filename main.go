package main

import (
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"encoding/json"
	// Routing library
	"github.com/gorilla/mux"
)

// Category structure
type Category struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

// Supplier structure
type Supplier struct {
	ID string `json:"id"`
	Name string `json:"name"`
	INN string `json:"inn"`
	CEO string `json:"ceo"`
	Phone string `json:"phone"`
}

// Product structure
type Product struct {
	ID string `json:"id"`
	Supplier *Supplier `json:"supplier"`
	Name string `json:"name"`
	Material string `json:"material"`
	Category *Category `json:"category"`
	Cost string `json:"cost"`
	Count string `json:"count"`
}

// Declaring an array of objects
var products []Product

// Get all products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Get product
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

// Create product
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(1000000))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)	
}

// Update product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(products)
}

// Delete product
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}

// Point of entry
func main() {
	r := mux.NewRouter()
	products = append(products, Product{
		ID:"1",
		Supplier: &Supplier{
			ID:"1",
			Name:"Мебельная фабрика",
			INN:"11111111111",
			CEO:"Иванов Иван Иванович",
			Phone: "98712312312",
		},
		Name:"Шкаф",
		Material:"Дуб",
		Category: &Category{
			ID:"1",
			Name:"Мебель",
		},
		Cost:"2990",
		Count:"5",
	})
	products = append(products, Product{
		ID:"2",
		Supplier: &Supplier{
			ID:"2",
			Name:"Техническая фабрика",
			INN:"2222222222",
			CEO:"Силиванов Сергей Сергеевич",
			Phone: "88712312312",
		},
		Name:"Смартфон",
		Material:"Полимеры",
		Category: &Category{
			ID:"2",
			Name:"Техника",
		},
		Cost:"8990",
		Count:"25",
	})
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}