package handlers

import (
	"kasir-api/models"
	"kasir-api/storage"
	"kasir-api/utils"
	"net/http"
)

// GET ALL
func GetCategories(w http.ResponseWriter, r *http.Request) {
	utils.EncodeJSON(w, http.StatusOK, storage.Categories)
}

// GET BY ID
func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r.URL.Path, "/api/category/")
	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	index := findCategoryById(id)
	if index < 0 {
		utils.EncodeJSON(w, http.StatusNotFound, map[string]string{"error": "category not found"})
		return
	}
	utils.EncodeJSON(w, http.StatusOK, storage.Categories[index])
}

// CREATE
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	if !utils.DecodeJSON(r, &category, w) {
		return
	}

	category.ID = storage.NextCategoryID()
	storage.Categories = append(storage.Categories, category)
	utils.EncodeJSON(w, http.StatusCreated, category)
}

// UPDATE
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r.URL.Path, "/api/category/")
	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	index := findCategoryById(id)
	if index < 0 {
		utils.EncodeJSON(w, http.StatusNotFound, map[string]string{"error": "category not found"})
		return
	}

	var category models.Category
	if !utils.DecodeJSON(r, &category, w) {
		return
	}

	category.ID = id
	storage.Categories[index] = category
	utils.EncodeJSON(w, http.StatusOK, category)
}

// DELETE
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r.URL.Path, "/api/category/")
	if err != nil {
		utils.EncodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	index := findCategoryById(id)
	if index < 0 {
		utils.EncodeJSON(w, http.StatusNotFound, map[string]string{"error": "category not found"})
		return
	}

	storage.Categories = append(storage.Categories[:index], storage.Categories[index+1:]...)
	utils.EncodeJSON(w, http.StatusOK, map[string]string{"message": "category deleted successfully"})
}

func findCategoryById(id int) int {
	for i, val := range storage.Categories {
		if val.ID == id {
			return i
		}
	}
	return -1
}
