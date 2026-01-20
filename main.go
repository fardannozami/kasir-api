package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var product = []Product{
	{
		ID:    2,
		Name:  "mie goreng",
		Price: 2000,
		Stock: 50,
	},
	{
		ID:    7,
		Name:  "mie ayam",
		Price: 3000,
		Stock: 100,
	},
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}

	for _, p := range product {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("product not found"))
	}
}
func updateProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}

	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid data"))
		return
	}

	for i, p := range product {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			updateProduct.ID = id
			product[i] = updateProduct
			json.NewEncoder(w).Encode(updateProduct)
			return
		}
	}
}

func deleteProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}

	for i, p := range product {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			product = append(product[:i], product[i+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("product delete successfully"))
			return
		}
	}
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "API Running",
		})
		w.Write([]byte("ok"))
	})

	// GET detail product http://localhost:8080/api/product/{id}
	// PUT update product http://localhost:8080/api/product/{id}
	// DELETE delete product http://localhost:8080/api/product/{id}
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProductById(w, r)
		case "PUT":
			updateProductById(w, r)
		case "DELETE":
			deleteProductById(w, r)
		}
	})

	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
		case "POST":
			// read data from request
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				fmt.Println("server gagal running")
			}

			// insert into product variable
			newProduct.ID = len(product) + 1
			product = append(product, newProduct)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
		}
	})

	fmt.Println("server running at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server gagal running")
	}

}
