package controllers

import (
	"backend/api"
	metrics "backend/infrastructure/metrics"
	"backend/repositories"
	"backend/services"
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type ProductController struct {
	DB                *gorm.DB
	productRepository repositories.ProductRepository
	productService    services.ProductService
}

func NewProductController(db *gorm.DB, productRepo repositories.ProductRepository, productService services.ProductService) *ProductController {
	return &ProductController{db, productRepo, productService}
}

func (pc *ProductController) GetProductBySlug(c *gin.Context) {
	productSlug := c.Param("slug")

	product, err := pc.productRepository.GetProductBySlug(productSlug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sentry.CaptureException(err)
			metrics.Error404Counter.WithLabelValues("404").Inc()
			api.SendError(c, http.StatusNotFound, "Продукт не найден")
			return
		}
		sentry.CaptureException(err)
		metrics.Error500Counter.WithLabelValues("500").Inc()
		api.SendError(c, http.StatusInternalServerError, "Не удалось получить данные о товаре")
		return
	}
	var photosResponse []api.PhotoResponse
	photosResponse = pc.productService.ConvertPhotosToResponse(product)

	var similarProducts []api.SimilarProductResponse
	similarProducts, err = pc.productRepository.GetSimilarProducts(product.CategoryId, product.ID)
	if err != nil {
		sentry.CaptureException(err)
		metrics.Error500Counter.WithLabelValues("500").Inc()
		api.SendError(c, http.StatusInternalServerError, "Не удалось получить Другие товары")
		return
	}

	response := api.ProductResponse{
		ID:              product.ID,
		Title:           product.Title,
		Description:     product.Description,
		Price:           product.Price,
		Category:        product.CategoryId,
		Slug:            product.Slug,
		Photos:          photosResponse,
		SimilarProducts: similarProducts,
	}

	metrics.SuccessCounter.WithLabelValues("200").Inc()
	c.JSON(http.StatusOK, response)
}

func (pc *ProductController) GetProductsByCategory(c *gin.Context) {
	categorySlug := c.Param("slug")

	var productsByCat api.GetProductsByCategoryResponse
	limit, err := pc.productService.GetLimitParam(c)
	if err != nil || limit < 1 {
		productsByCat, err = pc.productRepository.GetProductsByCategory(categorySlug, 0)
	} else {
		productsByCat, err = pc.productRepository.GetProductsByCategory(categorySlug, limit)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sentry.CaptureException(err)
			metrics.Error404Counter.WithLabelValues("404").Inc()
			api.SendError(c, http.StatusNotFound, "Категория не найдена")
			return
		}
		sentry.CaptureException(err)
		metrics.Error500Counter.WithLabelValues("500").Inc()
		api.SendError(c, http.StatusInternalServerError, "Ошибка при получении данных о товарах")
		return
	}

	metrics.SuccessCounter.WithLabelValues("200").Inc()
	c.JSON(http.StatusOK, productsByCat)
}
