package dto

import "github.com/imnmania/go_fiber_api/models"

type ProductDTO struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(product models.Product) ProductDTO {
	return ProductDTO{
		ID:           product.ID,
		Name:         product.Name,
		SerialNumber: product.SerialNumber,
	}
}
