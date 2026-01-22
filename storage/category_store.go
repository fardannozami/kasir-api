package storage

import "kasir-api/models"

// in memory repository
var Categories = []models.Category{
	{ID: 1, Name: "food", Description: "makanan"},
	{ID: 2, Name: "drink", Description: "minuman"},
}

func NextCategoryID() int {
	return len(Products) + 1
}
