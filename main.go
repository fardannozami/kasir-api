package main

import (
	"kasir-api/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handler.Health)

	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetProducts(w, r)
		case http.MethodPost:
			handler.CreateProduct(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetProductById(w, r)
		case http.MethodPut:
			handler.UpdateProduct(w, r)
		case http.MethodDelete:
			handler.DeleteProduct(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/category", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCategories(w, r)
		case http.MethodPost:
			handler.CreateCategory(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/category/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCategoryById(w, r)
		case http.MethodPut:
			handler.UpdateCategory(w, r)
		case http.MethodDelete:
			handler.DeleteCategory(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("🚀 server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
