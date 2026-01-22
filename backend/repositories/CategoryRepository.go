package repositories

import (
	"backend/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategory(categoryID int) (models.Category, error)
}
type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (ps *CategoryRepositoryImpl) GetCategory(categoryID int) (models.Category, error) {
	var category models.Category

	if err := ps.DB.First(&category, categoryID).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}
