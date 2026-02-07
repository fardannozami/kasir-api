package dto

import (
	"kasir-api/models"
	"time"
)

type TransactionResponse struct {
	ID          int                         `json:"id"`
	TotalAmount int                         `json:"total_amount"`
	CreatedAt   time.Time                   `json:"created_at"`
	Details     []TransactionDetailResponse `json:"details"`
}

type TransactionDetailResponse struct {
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name,omitempty"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}

func TransactionModelToResponse(transaction *models.Transaction) TransactionResponse {
	details := make([]TransactionDetailResponse, len(transaction.Details))

	for i, detail := range transaction.Details {
		details[i] = TransactionDetailResponse{
			TransactionID: transaction.ID,
			ProductID:     detail.ProductID,
			ProductName:   detail.ProductName,
			Quantity:      detail.Quantity,
			Subtotal:      detail.Subtotal,
		}
	}

	return TransactionResponse{
		ID:          transaction.ID,
		TotalAmount: transaction.TotalAmount,
		CreatedAt:   transaction.CreatedAt,
		Details:     details,
	}
}
