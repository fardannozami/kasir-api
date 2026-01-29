package dto

import "kasir-api/models"

type ProductRequest struct {
	Name       string `json:"name" validate:"required"`
	Price      int    `json:"price" validate:"required"`
	Stock      int    `json:"stock" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}

type ProductResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

func ProductToProductResponse(product *models.Product) ProductResponse {
	return ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.Category.ID,
	}
}

func ProductRequestToProduct(productRequest ProductRequest) models.Product {
	return models.Product{
		Name:  productRequest.Name,
		Price: productRequest.Price,
		Stock: productRequest.Stock,
		Category: models.Category{
			ID: productRequest.CategoryID,
		},
	}
}
