package repositories

import (
	"backend/dto"
	"backend/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (ps *ProductRepository) GetProductsByCategory(categoryID int64, limit int) ([]dto.ProductsSectionResponse, error) {
	var products []dto.ProductsSectionResponse

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

func (ps *ProductRepository) GetSimilarProducts(categoryID int, productID int) ([]dto.SimilarProductResponse, error) {
	var similarProducts []dto.SimilarProductResponse

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

func (ps *ProductRepository) GetProduct(productID int) (*models.Product, error) {
	var product models.Product

	if err := ps.DB.Preload("Photos").First(&product, productID).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
