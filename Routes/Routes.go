package Routes

import (
	"proj-mido/stripe-gateway/Controllers"

	"github.com/labstack/echo/v4"
)

// SetupRoutes configures routes for the application
func SetupRoutes(e *echo.Echo) {
	mido := e.Group("/mido")
	{
		mido.GET("/products", Controllers.GetProducts)
		mido.POST("/products", Controllers.CreateProducts)
		mido.GET("/config", Controllers.Config)
		mido.POST("/create-payment-intent", Controllers.HandleCreatePaymentIntent)
		mido.DELETE("/delete-product/:id", Controllers.DeleteProduct)
		// Uncomment and implement these if needed
		// mido.GET("/products/:id", Controllers.GetProductsByID)
		// mido.PUT("/products/:id", Controllers.UpdateProducts)
		// mido.DELETE("/products/:id", Controllers.DeleteProducts)
	}
}
