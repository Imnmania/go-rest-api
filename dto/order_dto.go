package dto

import (
	"time"

	"github.com/imnmania/go_fiber_api/models"
)

type OrderDTO struct {
	ID        uint        `json:"id"`
	User      *UserDTO    `json:"user,omitempty"`
	Product   *ProductDTO `json:"product,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
}

func CreateResponseOrder(order models.Order, userDTO *UserDTO, productDTO *ProductDTO) OrderDTO {
	return OrderDTO{
		ID:        order.ID,
		User:      userDTO,
		Product:   productDTO,
		CreatedAt: order.CreatedAt,
	}
}
