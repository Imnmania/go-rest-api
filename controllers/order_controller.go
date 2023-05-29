package controllers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/imnmania/go_fiber_api/config"
	"github.com/imnmania/go_fiber_api/dto"
	"github.com/imnmania/go_fiber_api/models"
)

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to parse request data",
		})
	}

	var user models.User
	if err := findUser(int(order.UserID), &user, c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "invalid user",
		})
	}

	var product models.Product
	if err := findProduct(int(order.ProductID), &product, c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "invalid product",
		})
	}

	if err := config.DB.Create(&order).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to create order",
		})
	}

	responseUser := dto.CreateResponseUser(user)
	responseProduct := dto.CreateResponseProduct(product)

	return c.Status(http.StatusCreated).JSON(dto.GenericResponse{
		Status: "success",
		Data:   dto.CreateResponseOrder(order, &responseUser, &responseProduct),
	})
}

func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order
	if err := config.DB.Find(&orders).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "could not fetch orders",
		})
	}

	var responseOrders []dto.OrderDTO
	for _, order := range orders {
		var user models.User
		findUser(int(order.UserID), &user, c)

		var product models.Product
		findProduct(int(order.ProductID), &product, c)

		responseUser := dto.CreateResponseUser(user)
		responseProduct := dto.CreateResponseProduct(product)
		responseOrder := dto.CreateResponseOrder(
			order,
			&responseUser,
			&responseProduct,
		)
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status: "success",
		Data:   responseOrders,
	})
}

func GetOrderByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	var order models.Order
	if err := findOrder(id, &order, c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	var user models.User
	findUser(int(order.UserID), &user, c)

	var product models.Product
	findProduct(int(order.ProductID), &product, c)

	responseUser := dto.CreateResponseUser(user)
	responseProduct := dto.CreateResponseProduct(product)
	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status: "success",
		Data: dto.CreateResponseOrder(
			order,
			&responseUser,
			&responseProduct,
		),
	})
}

// ----------------
// Helper functions
// ----------------
func findOrder(id int, order *models.Order, c *fiber.Ctx) error {
	config.DB.First(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order not found")
	}
	return nil
}
