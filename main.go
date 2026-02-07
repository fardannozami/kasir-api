package main

import (
	"kasir-api/config"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
	}

	config := config.Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	// run migration
	database.RunMigration(db)

	defer db.Close()

	// setup routes
	productRepository := repositories.NewProductRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	productService := services.NewProductService(productRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	http.HandleFunc("/", handlers.SwaggerUI)
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/swagger.json", handlers.SwaggerSpec)
	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)
	http.HandleFunc("/api/category", categoryHandler.HandleCategorys)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST

	log.Println("🚀 server running at", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
