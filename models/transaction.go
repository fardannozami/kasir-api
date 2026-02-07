package models

import "time"

type Transaction struct {
	ID          int                 `json:"id"`
	TotalAmount int                 `json:"total_amount"`
	CreatedAt   time.Time           `json:"created_at"`
	Details     []TransactionDetail `json:"details"`
}

type TransactionDetail struct {
	ID            int `json:"id"`
	TransactionID int `json:"transaction_id"`
	ProductID     int `json:"product_id"`
	Quantity      int `json:"quantity"`
	Subtotal      int `json:"subtotal"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type ReportData struct {
	TotalRevenue   int
	TotalTransaksi int
	ProdukTerlaris struct {
		Nama       string
		QtyTerjual int
	}
}
