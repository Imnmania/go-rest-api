package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/imnmania/go_fiber_api/config"
	"github.com/imnmania/go_fiber_api/controllers"
)

func init() {
	config.LoadEnvironmentVariables()
	config.ConnectDB()
}

func main() {
	log.Println("Welcome to REST API with Golang...")
	// General config
	app := fiber.New()

	// Routes
	setupRoutes(app)

	// Run app
	log.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoints
	app.Get("/", controllers.Welcome)

	// Common middlewares
	app.Use(recover.New()) // auto recovery
	app.Use(logger.New())  // auto logging

	// User endpoints
	userRoute := app.Group("/api/users")
	userRoute.Post("/", controllers.CreateUser)
	userRoute.Get("/", controllers.GetUsers)
	userRoute.Get("/:id", controllers.GetUserById)
	userRoute.Put("/:id", controllers.UpdateUser)
	userRoute.Delete("/:id", controllers.DeleteUser)
	userRoute.Delete("/", controllers.DeleteAllUsers)

	// Product endpoints
	productRoute := app.Group("/api/products")
	productRoute.Post("/", controllers.CreateProduct)
	productRoute.Get("/", controllers.GetProducts)
	productRoute.Get("/:id", controllers.GetProductById)
	productRoute.Put("/:id", controllers.UpdateProduct)
	productRoute.Delete("/:id", controllers.DeleteProductByID)
	productRoute.Delete("/", controllers.DeleteAllProducts)

	// Order endpoints
	orderRoute := app.Group("/api/orders")
	orderRoute.Post("/", controllers.CreateOrder)
	orderRoute.Get("/", controllers.GetOrders)
	orderRoute.Get("/:id", controllers.GetOrderByID)
}
