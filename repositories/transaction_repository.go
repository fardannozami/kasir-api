package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(
	items []models.CheckoutItem,
) (*models.Transaction, error) {

	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	// 1️⃣ Ambil produk + update stock
	for _, item := range items {
		var productName string
		var productPrice, stock int

		err := tx.QueryRow(`
			SELECT name, price, stock
			FROM products
			WHERE id = $1
			FOR UPDATE
		`, item.ProductID).Scan(&productName, &productPrice, &stock)

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("stock product %s not enough", productName)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec(`
			UPDATE products
			SET stock = stock - $1
			WHERE id = $2
		`, item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Subtotal:  subtotal,
		})
	}

	// 2️⃣ Insert transaction
	var transactionID int
	err = tx.QueryRow(`
		INSERT INTO transactions (total_amount)
		VALUES ($1)
		RETURNING id
	`, totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// 3️⃣ BULK INSERT transaction_details
	bulkQuery, args := buildBulkInsertDetails(details, transactionID)

	_, err = tx.Exec(bulkQuery, args...)
	if err != nil {
		return nil, err
	}

	// 4️⃣ Commit
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func buildBulkInsertDetails(
	details []models.TransactionDetail,
	transactionID int,
) (string, []interface{}) {

	query := `
		INSERT INTO transaction_details
		(transaction_id, product_id, quantity, subtotal)
		VALUES
	`

	args := []interface{}{}
	placeholder := 1

	for i, d := range details {
		query += fmt.Sprintf(
			"($%d, $%d, $%d, $%d)",
			placeholder,
			placeholder+1,
			placeholder+2,
			placeholder+3,
		)

		if i < len(details)-1 {
			query += ","
		}

		args = append(args,
			transactionID,
			d.ProductID,
			d.Quantity,
			d.Subtotal,
		)

		placeholder += 4
	}

	return query, args
}

func (repo *TransactionRepository) GetReportData(startDate, endDate string) (*models.ReportData, error) {
	var report models.ReportData

	// 1. Get Total Revenue and Total Transaksi
	err := repo.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0), COUNT(id)
		FROM transactions
		WHERE created_at::date BETWEEN $1 AND $2
	`, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	if report.TotalTransaksi == 0 {
		return nil, fmt.Errorf("data tidak ditemukan")
	}

	// 2. Get Produk Terlaris
	err = repo.db.QueryRow(`
		SELECT p.name, SUM(td.quantity) as qty
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE t.created_at::date BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY qty DESC
		LIMIT 1
	`, startDate, endDate).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)

	if err == sql.ErrNoRows {
		report.ProdukTerlaris.Nama = "-"
		report.ProdukTerlaris.QtyTerjual = 0
	} else if err != nil {
		return nil, err
	}

	return &report, nil
}
