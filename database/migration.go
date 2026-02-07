package database

import (
	"database/sql"
	"log"
)

func RunMigration(db *sql.DB) {
	// Create categories table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT
		);
	`)
	if err != nil {
		log.Fatal("Failed to create categories table:", err)
	}

	// Create products table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			price INT NOT NULL,
			stock INT NOT NULL,
			category_id INT,
			CONSTRAINT fk_category
				FOREIGN KEY(category_id)
				REFERENCES categories(id)
				ON DELETE SET NULL
		);
	`)
	if err != nil {
		log.Fatal("Failed to create products table:", err)
	}

	// Seed default category if not exists
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM categories WHERE id = 1").Scan(&count)
	if err != nil {
		log.Fatal("Failed to check categories:", err)
	}

	if count == 0 {
		_, err = db.Exec(`
			INSERT INTO categories (id, name, description) 
			VALUES (1, 'Sembako', 'Produk kebutuhan sehari-hari')
		`)
		if err != nil {
			log.Fatal("Failed to seed category:", err)
		}
		log.Println("Seeded default category")
	}

	// Create transactions table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			total_amount INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal("Failed to create transactions table:", err)
	}

	// Create transaction_details table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transaction_details (
			id SERIAL PRIMARY KEY,
			transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
			product_name VARCHAR(255) NOT NULL,
			price INT NOT NULL,
			quantity INT NOT NULL,
			category VARCHAR(255) NOT NULL,
			sub_total INT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("Failed to create transaction_details table:", err)
	}

	log.Println("Migration successful")
}
