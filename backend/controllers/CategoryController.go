package controllers

import (
	"backend/api"
	"backend/infrastructure/metrics"
	"backend/repositories"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB                 *gorm.DB
	categoryRepository repositories.CategoryRepository
}

func NewCategoryController(db *gorm.DB, categoryRepo repositories.CategoryRepository) *CategoryController {
	return &CategoryController{db, categoryRepo}
}

func (cc *CategoryController) GetCategoryBySlug(c *gin.Context) {
	categorySlug := c.Param("slug")

	category, err := cc.categoryRepository.GetCategoryBySlug(categorySlug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			metrics.Error404Counter.WithLabelValues("404").Inc()
			api.SendError(c, http.StatusNotFound, "Категория не найдена")
			return
		}
		metrics.Error500Counter.WithLabelValues("500").Inc()
		api.SendError(c, http.StatusInternalServerError, "Не удалось получить данные о категории")
		return
	}

	metrics.SuccessCounter.WithLabelValues("200").Inc()
	c.JSON(http.StatusOK, category)
}

func (cc *CategoryController) GetCategoriesList(c *gin.Context) {
	categories, err := cc.categoryRepository.GetAllCategories()

	if err != nil {
		metrics.Error500Counter.WithLabelValues("500").Inc()
		api.SendError(c, http.StatusInternalServerError, "Не удалось получить данные о категориях")
		return
	}

	metrics.SuccessCounter.WithLabelValues("200").Inc()
	c.JSON(http.StatusOK, categories)
}
