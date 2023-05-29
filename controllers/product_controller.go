package controllers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/imnmania/go_fiber_api/config"
	"github.com/imnmania/go_fiber_api/dto"
	"github.com/imnmania/go_fiber_api/models"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to parse request data",
		})
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to create product",
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.GenericResponse{
		Status: "success",
		Data:   dto.CreateResponseProduct(product),
	})
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "could not fetch products",
		})
	}

	var responseProducts []dto.ProductDTO
	for _, product := range products {
		responseProducts = append(responseProducts, dto.CreateResponseProduct(product))
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status: "success",
		Data:   responseProducts,
	})
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	var product models.Product
	if err := findProduct(id, &product, c); err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status: "success",
		Data:   dto.CreateResponseProduct(product),
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	var productToUpdate models.Product
	if err := findProduct(id, &productToUpdate, c); err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	var responseBody struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}
	if err := c.BodyParser(&responseBody); err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "error in request body",
		})
	}

	productToUpdate.Name = responseBody.Name
	productToUpdate.SerialNumber = responseBody.SerialNumber

	if err := config.DB.Save(&productToUpdate).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to update product",
		})
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status: "success",
		Data:   dto.CreateResponseProduct(productToUpdate),
	})
}

func DeleteProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	var productToDelete models.Product
	if err := findProduct(id, &productToDelete, c); err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	if err := config.DB.Delete(&productToDelete).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to delete product",
		})
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status:  "success",
		Message: "successfully deleted the product",
	})
}

func DeleteAllProducts(c *fiber.Ctx) error {
	if err := config.DB.Exec("DELETE FROM products").Error; err != nil {
		return c.Status(http.StatusConflict).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to delete all products",
		})
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status:  "success",
		Message: "successfully deleted all products",
	})
}

// ----------------
// Helper functions
// ----------------
func findProduct(id int, product *models.Product, c *fiber.Ctx) error {
	config.DB.First(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product not found")
	}
	return nil
}
