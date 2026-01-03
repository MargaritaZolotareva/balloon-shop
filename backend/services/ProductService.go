package services

import (
	"backend/dto"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProductService struct{}

func NewProductService() ProductService {
	return ProductService{}
}

func (ps *ProductService) ValidateCategoryID(categoryID string) (int64, error) {
	if categoryID == "" {
		return 0, fmt.Errorf("Необходимо указать category_id")
	}

	categoryInt, err := strconv.ParseInt(categoryID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Неверный формат category_id")
	}

	return categoryInt, nil
}

func (ps *ProductService) ConvertPhotosToResponse(product *models.Product) []dto.PhotoResponse {
	var photosResponse []dto.PhotoResponse
	for _, p := range product.Photos {
		photosResponse = append(photosResponse, dto.PhotoResponse{
			ID:        p.ID,
			PhotoPath: p.PhotoPath,
		})
	}

	return photosResponse
}

func (ps *ProductService) GetLimitParam(c *gin.Context) (int, error) {
	limit := c.DefaultQuery("limit", "0")
	return strconv.Atoi(limit)
}
