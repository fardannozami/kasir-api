package handlers

import (
	"encoding/json"
	"kasir-api/dto"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"time"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// multiple item apa aja, quantity nya
func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	case http.MethodGet:
		h.HandleReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.TransactionModelToResponse(transaction)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TransactionHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	// If it's the "hari-ini" endpoint or dates are missing, use today
	if r.URL.Path == "/api/report/hari-ini" || (startDate == "" && endDate == "") {
		today := time.Now().Format("2006-01-02")
		startDate = today
		endDate = today
	}

	reportData, err := h.service.GetReport(startDate, endDate)
	if err != nil {
		if err.Error() == "data tidak ditemukan" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ReportResponse{
		TotalRevenue:   reportData.TotalRevenue,
		TotalTransaksi: reportData.TotalTransaksi,
		ProdukTerlaris: dto.ProdukTerlarisResponse{
			Nama:       reportData.ProdukTerlaris.Nama,
			QtyTerjual: reportData.ProdukTerlaris.QtyTerjual,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
