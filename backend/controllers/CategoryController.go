package controllers

import (
	"backend/models"
	"backend/repositories"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID категории"})
		return
	}

	var category models.Category
	category, err = cc.categoryRepository.GetCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные о категории"})
		return
	}

	c.JSON(http.StatusOK, category)
}
