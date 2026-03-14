package repositories

import (
	"backend/api"
	"backend/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoryBySlug(categorySlug string) (models.Category, error)
	GetAllCategories() ([]api.CategoriesResponse, error)
	GetHomepageCategories() (api.HomepageResponse, error)
}
type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (ps *CategoryRepositoryImpl) GetCategoryBySlug(categorySlug string) (models.Category, error) {
	var category models.Category

	if err := ps.DB.First(&category, categorySlug).Error; err != nil {
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

func (ps *CategoryRepositoryImpl) GetHomepageCategories() (api.HomepageResponse, error) {
	type row struct {
		CategoryID    int     `gorm:"column:category_id"`
		CategoryTitle string  `gorm:"column:category_title"`
		CategorySlug  string  `gorm:"column:category_slug"`
		ProductID     int     `gorm:"column:product_id"`
		ProductTitle  string  `gorm:"column:product_title"`
		Price         float64 `gorm:"column:price"`
		ProductSlug   string  `gorm:"column:product_slug"`
		Photo         string  `gorm:"column:photo"`
	}

	var rows []row
	query := "SELECT c.id AS category_id, c.title AS category_title, c.slug AS category_slug, p.id AS product_id, p.title AS product_title, p.price, p.slug AS product_slug, ph.photo_path AS photo " +
		"FROM categories c " +
		"LEFT JOIN (" +
		"SELECT *, ROW_NUMBER() OVER (PARTITION BY category_id ORDER BY product_order ASC) AS rn " +
		"FROM products) p " +
		"ON c.id = p.category_id AND p.rn <= 5 " +
		"LEFT JOIN photos ph ON p.id = ph.product_id AND ph.pic_order = 0 " +
		"WHERE c.category_order IS NOT NULL " +
		"ORDER BY c.category_order"

	if err := ps.DB.Raw(query).Scan(&rows).Error; err != nil {
		return api.HomepageResponse{}, err
	}

	categoriesMap := map[int]*api.HomepageCategory{}
	var categoriesList []api.HomepageCategory

	for _, r := range rows {
		cat, ok := categoriesMap[r.CategoryID]
		if !ok {
			cat = &api.HomepageCategory{
				Title:    r.CategoryTitle,
				Slug:     r.CategorySlug,
				Products: []api.ProductsSectionResponse{},
			}
			categoriesMap[r.CategoryID] = cat
			categoriesList = append(categoriesList, *cat)
			cat = &categoriesList[len(categoriesList)-1]
			categoriesMap[r.CategoryID] = cat
		}

		if r.ProductID != 0 {
			cat.Products = append(cat.Products, api.ProductsSectionResponse{
				ID:            r.ProductID,
				Title:         r.ProductTitle,
				Price:         r.Price,
				Photo:         r.Photo,
				Slug:          r.ProductSlug,
				CategoryTitle: r.CategoryTitle,
			})
		}
	}

	return api.HomepageResponse{Categories: categoriesList}, nil
}
