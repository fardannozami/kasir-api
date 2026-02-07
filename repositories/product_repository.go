package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(name string) ([]models.Product, error) {
	query := `
		SELECT
			products.id,
			products.name,
			products.price,
			products.stock,
			categories.id AS category_id,
			categories.name AS category_name
		FROM products
		JOIN categories ON products.category_id = categories.id
	`

	var args []interface{}
	if name != "" {
		query += " WHERE products.name ILIKE $1"
		args = append(args, "%"+name+"%")
	}

	query += " ORDER BY products.id DESC"

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.Stock,
			&p.Category.ID,
			&p.Category.Name,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	query := "INSERT INTO products (name, category_id, price, stock) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Category.ID, product.Price, product.Stock).Scan(&product.ID)
	return product, err
}

// GetByID - ambil product by ID
func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `
		SELECT
			products.id,
			products.name,
			products.price,
			products.stock,
			categories.id AS category_id,
			categories.name as category_name
		FROM products
		JOIN categories ON products.category_id = categories.id
		WHERE products.id = $1
	`

	var p models.Product
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category.ID, &p.Category.Name)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *ProductRepository) Update(product *models.Product) (*models.Product, error) {
	log.Printf("%+v\n", product)

	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.Category.ID, product.ID)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("produk tidak ditemukan")
	}

	return product, nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
