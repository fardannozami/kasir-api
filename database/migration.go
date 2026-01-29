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

	log.Println("Migration successful")
}
