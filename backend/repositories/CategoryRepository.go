package repositories

import (
	"backend/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepository{db}
}

func (ps *CategoryRepository) GetCategory(categoryID int) (models.Category, error) {
	var category models.Category

	if err := ps.DB.First(&category, categoryID).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}
