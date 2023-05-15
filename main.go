package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-ecommerce-api/controllers"
	"github.com/golang-ecommerce-api/initializers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.NewConnection()
	initializers.SyncDB()
}

func main() {
	app := fiber.New()

	app.Post("/signup", controllers.SignUp)
	app.Post("/login", controllers.Login)

	//use middleware here
	customers := app.Group("/customers")
	customers.Get("/", controllers.GetCustomers)
	customers.Get("/:id", controllers.GetCustomerByID)
	customers.Patch("/:id", controllers.UpdateCustomer)
	customers.Delete("/:id", controllers.DeleteCustomer)

	products := app.Group("/products")
	products.Get("/", controllers.GetProducts)
	products.Post("/", controllers.CreateProduct)
	products.Get("/:id", controllers.GetProductByID)
	products.Patch("/:id", controllers.UpdateProduct)
	products.Delete("/:id", controllers.DeleteProduct)

	app.Listen(":8000")
}
