package services

import (
	"backend/api"
	"backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	ConvertPhotosToResponse(product models.Product) []api.PhotoResponse
	GetLimitParam(c *gin.Context) (int, error)
}

type ProductServiceImpl struct{}

func NewProductService() ProductService {
	return &ProductServiceImpl{}
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
