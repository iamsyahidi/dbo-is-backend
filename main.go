package main

import (
	"context"
	"dbo/dbo-is-backend/controllers"
	"dbo/dbo-is-backend/pkg/database"
	"dbo/dbo-is-backend/repositories"
	"dbo/dbo-is-backend/routes"
	"dbo/dbo-is-backend/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title DBO Backend Test API
// @description This is a backend test API for DBO.
// @version 1
// @host localhost:3001
// @BasePath /v1
// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
// @contact.name Ilham Syahidi
// @contact.email ilhamsyahidi66@gmail.com
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := database.InitGorm(ctx)
	if err != nil {
		panic(err)
	}

	gin.ForceConsoleColor()
	gin.SetMode(gin.DebugMode)

	// Repositories
	customerRepository := repositories.NewCustomerRepository(db)
	productRepository := repositories.NewProductRepository(db)
	orderRepository := repositories.NewOrderRepository(db)

	// Services
	customerService := services.NewCustomerService(customerRepository)
	productService := services.NewProductService(productRepository)
	authService := services.NewAuthService(customerRepository)
	orderService := services.NewOrderService(orderRepository, productRepository)

	// Controllers
	customerController := controllers.NewCustomerController(customerService)
	productController := controllers.NewProductController(productService)
	authController := controllers.NewAuthController(authService)
	orderController := controllers.NewOrderController(orderService)

	router := routes.NewRouter(customerController, productController, authController, orderController)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	router.Run(port)
}

//TODO add swagger docs, setup env db
