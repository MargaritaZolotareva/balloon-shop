package controllers

import (
	"backend/api"
	"backend/infrastructure/metrics"
	"backend/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type HomepageController struct {
	DB                 *gorm.DB
	categoryRepository repositories.CategoryRepository
}

func NewHomepageController(db *gorm.DB, categoryRepo repositories.CategoryRepository) *HomepageController {
	return &HomepageController{db, categoryRepo}
}

func (hc *HomepageController) GetHomepageCategories(c *gin.Context) {
	categoriesData, err := hc.categoryRepository.GetHomepageCategories()

	if err != nil {
		metrics.Error500Counter.WithLabelValues("500").Inc()
		api.SendError(c, http.StatusInternalServerError, "Не удалось получить данные о категориях")
		return
	}

	metrics.SuccessCounter.WithLabelValues("200").Inc()
	c.JSON(http.StatusOK, categoriesData)
}
