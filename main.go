package main

import (
	"fmt"
	"proj-mido/stripe-gateway/Config"
	"proj-mido/stripe-gateway/Models"
	"proj-mido/stripe-gateway/Routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	// Initialize the database connection
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
		return
	}
	defer Config.DB.Close()

	// Auto-migrate the models
	Config.DB.AutoMigrate(&Models.Products{})

	// Create a new Echo instance
	e := echo.New()

	// Set up CORS middleware
	e.Use(middleware.CORS())

	// Set up routes
	Routes.SetupRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
