package Controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proj-mido/stripe-gateway/Models"
	"proj-mido/stripe-gateway/Repository"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

func GetProducts(c echo.Context) error {
	var products []Models.Products
	err := Repository.GetAllProducts(&products)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "Products not found"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"products": products})
}

func CreateProducts(c echo.Context) error {
	var product Models.Products
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid input"})
	}
	err := Repository.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Error creating product"})
	}
	return c.JSON(http.StatusOK, product)
}

func Config(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"publishableKey": os.Getenv("STRIPE_PUBLISHABLE_KEY"),
	})
}

// HandleCreatePaymentIntent handles creating a payment intent------------------->
func HandleCreatePaymentIntent(c echo.Context) error {
	var product Models.Products

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	productID := strconv.FormatUint(uint64(product.Id), 10)

	data, err := Repository.GetAProduct(&product, productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "Product not found"})
	}

	fmt.Println("Data==>", data)

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(data.Price)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	log.Printf("pi.New: %v", pi.ClientSecret)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"clientSecret": pi.ClientSecret,
	})
}







// code with clear responses, Aligned the above with reach rendering, THE CODE IS THE SAME

// package Controllers

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"proj-mido/stripe-gateway/Models"
// 	"proj-mido/stripe-gateway/Repository"
// 	"strconv"

// 	"github.com/joho/godotenv"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stripe/stripe-go/v72"
// 	"github.com/stripe/stripe-go/v72/paymentintent"
// )

// // StandardResponse represents a standard API response
// type StandardResponse struct {
// 	Status  bool        `json:"status"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data,omitempty"`
// }

// // GetProducts handles fetching all products
// func GetProducts(c echo.Context) error {
// 	var products []Models.Products
// 	err := Repository.GetAllProducts(&products)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, StandardResponse{false, "Products not found", nil})
// 	}
// 	return c.JSON(http.StatusOK, StandardResponse{true, "Products retrieved successfully", products})
// }

// // CreateProducts handles creating a new product
// func CreateProducts(c echo.Context) error {
// 	var product Models.Products
// 	if err := c.Bind(&product); err != nil {
// 		return c.JSON(http.StatusBadRequest, StandardResponse{false, "Invalid input", nil})
// 	}
// 	err := Repository.CreateProduct(&product)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, StandardResponse{false, "Error creating product", nil})
// 	}
// 	return c.JSON(http.StatusOK, StandardResponse{true, "Product created successfully", product})
// }

// // Config loads environment variables and returns the Stripe publishable key
// func Config(c echo.Context) error {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	return c.JSON(http.StatusOK, StandardResponse{true, "Config loaded successfully", map[string]interface{}{
// 		"publishableKey": os.Getenv("STRIPE_PUBLISHABLE_KEY"),
// 	}})
// }

// // HandleCreatePaymentIntent handles creating a payment intent
// func HandleCreatePaymentIntent(c echo.Context) error {
// 	var product Models.Products

// 	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

// 	if err := c.Bind(&product); err != nil {
// 		return c.JSON(http.StatusInternalServerError, StandardResponse{false, err.Error(), nil})
// 	}

// 	productID := strconv.FormatUint(uint64(product.Id), 10)

// 	data, err := Repository.GetAProduct(&product, productID)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, StandardResponse{false, "Product not found", nil})
// 	}

// 	fmt.Println("Data==>", data)

// 	// Create a PaymentIntent with amount and currency
// 	params := &stripe.PaymentIntentParams{
// 		Amount:   stripe.Int64(int64(data.Price)),
// 		Currency: stripe.String(string(stripe.CurrencyUSD)),
// 		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
// 			Enabled: stripe.Bool(true),
// 		},
// 	}

// 	pi, err := paymentintent.New(params)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, StandardResponse{false, err.Error(), nil})
// 	}
// 	log.Printf("pi.New: %v", pi.ClientSecret)

// 	return c.JSON(http.StatusOK, StandardResponse{true, "PaymentIntent created successfully", map[string]interface{}{
// 		"clientSecret": pi.ClientSecret,
// 	}})
// }
