package dto

import "kasir-api/models"

type CategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CategoryToCategoryResponse(category *models.Category) CategoryResponse {
	return CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func CategoryRequestToCategory(categoryRequest CategoryRequest) models.Category {
	return models.Category{
		Name:        categoryRequest.Name,
		Description: categoryRequest.Description,
	}
}
