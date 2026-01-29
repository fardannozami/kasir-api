package handler

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

	defer db.Close()

	// setup routes
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)
	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	log.Println("🚀 server running at", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
