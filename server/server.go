package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
	"productsServer/database"
	"productsServer/models"
	"productsServer/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	// Ważne, bez tego nie sprawdza kluczy
	db.Exec("PRAGMA foreign_keys = ON")

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Product{}, &models.CartItem{}, &models.Cart{}, &models.Payment{})
	if err != nil {
		fmt.Println("Failed to migrate database")
		return
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	allowedFrontend := os.Getenv("FRONTEND_URL")

	e := echo.New()
	database.InitDB(db)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowOriginFunc: func(origin string) (bool, error) {
			return origin == allowedFrontend, nil
		},
	}))

	routers.RegisterProductRoutes(e)
	routers.RegisterPaymentCartRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
