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
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
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
