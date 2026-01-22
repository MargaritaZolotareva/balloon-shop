package repositories

import (
	"backend/api"
	"backend/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductsByCategory(categoryID int, limit int) ([]api.ProductsSectionResponse, error)
	GetSimilarProducts(categoryID int, productID int) ([]api.SimilarProductResponse, error)
	GetProduct(productID int) (models.Product, error)
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {

	return &ProductRepositoryImpl{DB: db}
}

func (ps *ProductRepositoryImpl) GetProductsByCategory(categoryID int, limit int) ([]api.ProductsSectionResponse, error) {
	var products []api.ProductsSectionResponse

	query := ps.DB.Table("products").
		Select("products.id, products.title, products.price, photos.photo_path as photo").
		Joins("LEFT JOIN photos ON photos.product_id = products.id").
		Where("products.category_id = ?", categoryID).
		Order("photos.pic_order ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductRepositoryImpl) GetSimilarProducts(categoryID int, productID int) ([]api.SimilarProductResponse, error) {
	var similarProducts []api.SimilarProductResponse

	query := ps.DB.Table("products").
		Select("products.id, products.title, products.price, photos.photo_path as photo").
		Joins("LEFT JOIN photos ON photos.product_id = products.id").
		Where("category_id = ? AND products.id <> ?", categoryID, productID).
		Order("photos.pic_order ASC").
		Limit(10)

	if err := query.Find(&similarProducts).Error; err != nil {
		return nil, err
	}

	return similarProducts, nil
}

func (ps *ProductRepositoryImpl) GetProduct(productID int) (models.Product, error) {
	var product models.Product

	if err := ps.DB.Preload("Photos").First(&product, productID).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}
