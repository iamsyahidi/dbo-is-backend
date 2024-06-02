package routes

import (
	"dbo/dbo-is-backend/controllers"
	"dbo/dbo-is-backend/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(customerController controllers.CustomerControllerInterface, productController controllers.ProductControllerInterface, authController controllers.AuthControllerInterface, orderController controllers.OrderControllerInterface) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	baseRouter := router.Group("/v1")

	//* customers
	customers := baseRouter.Group("/customers")
	customers.POST("", customerController.CreateCustomer)
	customersWithAuth := baseRouter.Group("/customers")
	customersWithAuth.Use(middleware.AuthMiddleware())
	customersWithAuth.GET("", customerController.GetCustomers)
	customersWithAuth.GET("/:id", customerController.GetCustomerById)
	customersWithAuth.PUT("/:id", customerController.UpdateCustomer)
	customersWithAuth.DELETE("/:id", customerController.DeleteCustomer)

	//* products
	products := baseRouter.Group("/products")
	products.POST("", productController.CreateProduct)
	productsWithAuth := baseRouter.Group("/products")
	productsWithAuth.Use(middleware.AuthMiddleware())
	productsWithAuth.GET("", productController.GetProducts)
	productsWithAuth.GET("/:id", productController.GetProductById)
	productsWithAuth.PUT("/:id", productController.UpdateProduct)
	productsWithAuth.DELETE("/:id", productController.DeleteProduct)

	//* auth
	auth := baseRouter.Group("/auth")
	auth.POST("/login", authController.Login)

	//* orders
	orders := baseRouter.Group("/orders")
	orders.Use(middleware.AuthMiddleware())
	orders.POST("", orderController.CreateOrder)

	return router
}
