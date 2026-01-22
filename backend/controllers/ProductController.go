package controllers

import (
	"backend/api"
	metrics "backend/infrastructure/metrics"
	"backend/models"
	"backend/repositories"
	"backend/services"
	"errors"
	"github.com/getsentry/sentry-go"
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
		sentry.CaptureException(err)
		metrics.Error400Counter.WithLabelValues("400").Inc()
		api.SendError(c, http.StatusBadRequest, "Неверный формат ID товара")
		return
	}

	var product models.Product
	product, err = pc.productRepository.GetProduct(productID)
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
	similarProducts, err = pc.productRepository.GetSimilarProducts(product.CategoryId, productID)
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
		Photos:          photosResponse,
		SimilarProducts: similarProducts,
	}

	metrics.SuccessCounter.WithLabelValues("200").Inc()
	c.JSON(http.StatusOK, response)
}

func (pc *ProductController) GetProductsByCategory(c *gin.Context) {
	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		sentry.CaptureException(err)
		metrics.Error400Counter.WithLabelValues("400").Inc()
		api.SendError(c, http.StatusBadRequest, "Неверный формат ID категории")
		return
	}

	var limit int
	var products []api.ProductsSectionResponse
	limit, err = pc.productService.GetLimitParam(c)
	if err != nil || limit < 1 {
		products, err = pc.productRepository.GetProductsByCategory(int(categoryID), 0)
	} else {
		products, err = pc.productRepository.GetProductsByCategory(int(categoryID), limit)
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
	c.JSON(http.StatusOK, products)
}
