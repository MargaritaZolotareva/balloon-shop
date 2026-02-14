package controllers

import (
	"backend/api"
	"backend/infrastructure/metrics"
	"backend/models"
	"backend/repositories"
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CategoryController struct {
	DB                 *gorm.DB
	categoryRepository repositories.CategoryRepository
}

func NewCategoryController(db *gorm.DB, categoryRepo repositories.CategoryRepository) *CategoryController {
	return &CategoryController{db, categoryRepo}
}

func (cc *CategoryController) GetCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sentry.CaptureException(err)
		metrics.Error400Counter.WithLabelValues("400").Inc()
		api.SendError(c, http.StatusBadRequest, "Неверный формат ID категории")
		return
	}

	var category models.Category
	category, err = cc.categoryRepository.GetCategory(categoryID)
	if err != nil {
		sentry.CaptureException(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			metrics.Error404Counter.WithLabelValues("404").Inc()
			api.SendError(c, http.StatusNotFound, "Товар не найден")
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
		sentry.CaptureException(err)
		metrics.Error500Counter.WithLabelValues("500").Inc()
		api.SendError(c, http.StatusInternalServerError, "Не удалось получить данные о категориях")
		return
	}

	metrics.SuccessCounter.WithLabelValues("200").Inc()
	c.JSON(http.StatusOK, categories)
}
