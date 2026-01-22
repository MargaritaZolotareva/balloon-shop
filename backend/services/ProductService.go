package services

import (
	"backend/api"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProductService interface {
	ValidateCategoryID(categoryID string) (int, error)
	ConvertPhotosToResponse(product models.Product) []api.PhotoResponse
	GetLimitParam(c *gin.Context) (int, error)
}

type ProductServiceImpl struct{}

func NewProductService() ProductService {
	return &ProductServiceImpl{}
}

func (ps *ProductServiceImpl) ValidateCategoryID(categoryID string) (int, error) {
	if categoryID == "" {
		return 0, fmt.Errorf("Необходимо указать category_id")
	}

	categoryInt, err := strconv.ParseInt(categoryID, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("Неверный формат category_id")
	}

	return int(categoryInt), nil
}

func (ps *ProductServiceImpl) ConvertPhotosToResponse(product models.Product) []api.PhotoResponse {
	var photosResponse []api.PhotoResponse
	for _, p := range product.Photos {
		photosResponse = append(photosResponse, api.PhotoResponse{
			ID:        p.ID,
			PhotoPath: p.PhotoPath,
		})
	}

	return photosResponse
}

func (ps *ProductServiceImpl) GetLimitParam(c *gin.Context) (int, error) {
	limit := c.DefaultQuery("limit", "0")
	return strconv.Atoi(limit)
}
