package handler

import (
	"kasir-api/models"
	"kasir-api/storage"
	"kasir-api/utils"
	"net/http"
)

// GET ALL
func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.EncodeJSON(w, http.StatusOK, storage.Products)
}

// GET BY ID
func GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r.URL.Path, "/api/product/")
	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	index := findProductById(id)
	if index < 0 {
		utils.EncodeJSON(w, http.StatusNotFound, map[string]string{"error": "product not found"})
		return
	}
	utils.EncodeJSON(w, http.StatusOK, storage.Products[index])
}

// CREATE
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if !utils.DecodeJSON(r, &product, w) {
		return
	}

	product.ID = storage.NextProductID()
	storage.Products = append(storage.Products, product)
	utils.EncodeJSON(w, http.StatusCreated, product)
}

// UPDATE
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r.URL.Path, "/api/product/")
	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	index := findProductById(id)
	if index < 0 {
		utils.EncodeJSON(w, http.StatusNotFound, map[string]string{"error": "product not found"})
		return
	}

	var product models.Product
	if !utils.DecodeJSON(r, &product, w) {
		return
	}

	product.ID = id
	storage.Products[index] = product
	utils.EncodeJSON(w, http.StatusOK, product)
}

// DELETE
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r.URL.Path, "/api/product/")
	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	index := findProductById(id)
	if index < 0 {
		utils.EncodeJSON(w, http.StatusNotFound, map[string]string{"error": "product not found"})
		return
	}

	storage.Products = append(storage.Products[:index], storage.Products[index+1:]...)
	utils.EncodeJSON(w, http.StatusOK, map[string]string{"message": "product deleted successfully"})
}

func findProductById(id int) int {
	for i, val := range storage.Products {
		if val.ID == id {
			return i
		}
	}
	return -1
}
