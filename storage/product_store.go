package storage

import "kasir-api/models"

// in memory repository
var Products = []models.Product{
	{ID: 1, Name: "mie goreng", Price: 2000, Stock: 50},
	{ID: 2, Name: "mie ayam", Price: 3000, Stock: 100},
}

func NextProductID() int {
	return len(Products) + 1
}
