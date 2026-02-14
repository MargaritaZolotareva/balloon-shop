package repositories

import (
	"backend/api"
	"backend/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategory(categoryID int) (models.Category, error)
	GetAllCategories() ([]api.CategoriesResponse, error)
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

func (ps *CategoryRepositoryImpl) GetAllCategories() ([]api.CategoriesResponse, error) {
	var categories []api.CategoriesResponse

	if err := ps.DB.Table("categories").
		Joins("LEFT JOIN photos ON photos.id = categories.photo_id").
		Select("categories.id, categories.title, photos.photo_path AS photo").
		Find(&categories).Error; err != nil {
		return []api.CategoriesResponse{}, err
	}

	return categories, nil
}
