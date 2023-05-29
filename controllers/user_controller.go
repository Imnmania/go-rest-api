package controllers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/imnmania/go_fiber_api/config"
	"github.com/imnmania/go_fiber_api/dto"
	"github.com/imnmania/go_fiber_api/models"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to parse body",
		})
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusConflict).JSON(dto.GenericResponse{
			Status:  "success",
			Message: "failed to create user",
		})
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status: "success",
		Data:   dto.CreateResponseUser(user),
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return c.Status(http.StatusConflict).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to fetch users",
		})
	}

	var userDTOs []dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, dto.CreateResponseUser(user))
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status: "success",
		Data:   userDTOs,
	})
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "invalid id provided for user",
		})
	}

	var user models.User
	if err := findUser(id, &user, c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status: "success",
		Data:   dto.CreateResponseUser(user),
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "success",
			Message: "invalid id provided for user",
		})
	}

	var user models.User
	if err := findUser(id, &user, c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	var requestBody struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "invalid request body",
		})
	}

	user.FirstName = requestBody.FirstName
	user.LastName = requestBody.LastName

	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: "failed to save user",
		})
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status: "success",
		Data:   dto.CreateResponseUser(user),
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "success",
			Message: "invalid id provided for user",
		})
	}

	var user models.User
	if err := findUser(id, &user, c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.GenericResponse{
			Status: "failed",
			Data:   "failed to delete user",
		})
	}

	return c.Status(http.StatusOK).JSON(dto.GenericResponse{
		Status:  "success",
		Message: "successfully deleted user",
	})
}

func DeleteAllUsers(c *fiber.Ctx) error {
	if err := config.DB.Exec("DELETE FROM users").Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to delete all user data",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "successfully deleted all user data",
	})
}

// ----------------
// Helper functions
// ----------------
func findUser(id int, user *models.User, c *fiber.Ctx) error {
	config.DB.First(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user not found")
	}
	return nil
}
