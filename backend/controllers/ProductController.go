package controllers

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ProductController struct {
	DB                *gorm.DB
	productRepository repositories.ProductRepository
	productService    services.ProductService
}

func NewProductController(db *gorm.DB, productRepo repositories.ProductRepository, productService services.ProductService) *ProductController {
	return &ProductController{db, productRepo, productService}
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID товара"})
		return
	}

	var product *models.Product
	product, err = pc.productRepository.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные о товаре"})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
		return
	}

	var photosResponse []dto.PhotoResponse
	photosResponse = pc.productService.ConvertPhotosToResponse(product)

	var similarProducts []dto.SimilarProductResponse
	similarProducts, err = pc.productRepository.GetSimilarProducts(product.CategoryId, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить Другие товары"})
		return
	}

	response := dto.ProductResponse{
		ID:              product.ID,
		Title:           product.Title,
		Description:     product.Description,
		Price:           product.Price,
		Category:        product.CategoryId,
		Photos:          photosResponse,
		SimilarProducts: similarProducts,
	}

	c.JSON(http.StatusOK, response)
}

func (pc *ProductController) GetProductsByCategory(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID категории"})
		return
	}

	var limit int
	var products []dto.ProductsSectionResponse
	limit, err = pc.productService.GetLimitParam(c)
	if err != nil || limit < 1 {
		products, err = pc.productRepository.GetProductsByCategory(categoryID, 0)
	} else {
		products, err = pc.productRepository.GetProductsByCategory(categoryID, limit)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных о товарах"})
		return
	}
	c.JSON(http.StatusOK, products)
}
