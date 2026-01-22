package controllers

import (
	"backend/api"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockProductRepo struct {
	mock.Mock
}
type MockProductService struct {
	mock.Mock
}

func (m *MockProductRepo) GetProduct(productID int) (models.Product, error) {
	args := m.Called(productID)
	return args.Get(0).(models.Product), args.Error(1)
}
func (m *MockProductRepo) GetSimilarProducts(categoryID, productID int) ([]api.SimilarProductResponse, error) {
	args := m.Called(categoryID, productID)
	return args.Get(0).([]api.SimilarProductResponse), args.Error(1)
}
func (m *MockProductRepo) GetProductsByCategory(categoryID int, limit int) ([]api.ProductsSectionResponse, error) {
	args := m.Called(categoryID, limit)
	return args.Get(0).([]api.ProductsSectionResponse), args.Error(1)
}

func (m *MockProductService) ValidateCategoryID(categoryID string) (int, error) {
	args := m.Called(categoryID)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockProductService) ConvertPhotosToResponse(product models.Product) []api.PhotoResponse {
	args := m.Called(product)
	return args.Get(0).([]api.PhotoResponse)
}

func (m *MockProductService) GetLimitParam(c *gin.Context) (int, error) {
	args := m.Called(c)
	return args.Get(0).(int), args.Error(1)
}

func TestProductController_GetProduct_Success(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)

	product := models.Product{
		ID:          1,
		Title:       "Test Product",
		Description: "Test description",
		Price:       100,
		CategoryId:  1,
	}

	mockRepo.On("GetProduct", product.ID).Return(product, nil)
	mockService.On("ConvertPhotosToResponse", product).Return([]api.PhotoResponse{
		{
			ID:        1,
			PhotoPath: "/photo1.jpg",
		},
		{
			ID:        2,
			PhotoPath: "/photo2.jpg",
		},
	})
	mockRepo.On("GetSimilarProducts", product.CategoryId, product.ID).Return([]api.SimilarProductResponse{
		{
			ID:    2,
			Title: "Similar Product 1",
			Price: 99.99,
			Photo: "/similar1.jpg",
		},
		{
			ID:    3,
			Title: "Similar Product 2",
			Price: 149.99,
			Photo: "/similar2.jpg",
		},
	}, nil)

	r := gin.Default()
	pc := NewProductController(nil, mockRepo, mockService)
	r.GET("/product/:id", pc.GetProduct)

	req, _ := http.NewRequest(http.MethodGet, "/product/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product")

	mockRepo.AssertExpectations(t)
	mockService.AssertExpectations(t)
}

func TestProductController_GetProduct_InvalidID(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)

	r := gin.Default()
	pc := NewProductController(nil, mockRepo, mockService)
	r.GET("/product/:id", pc.GetProduct)

	req, _ := http.NewRequest(http.MethodGet, "/product/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Неверный формат ID товара")
}

func TestProductController_GetProduct_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)

	mockRepo.On("GetProduct", 100500).Return(models.Product{}, gorm.ErrRecordNotFound)

	r := gin.Default()
	pc := NewProductController(nil, mockRepo, mockService)
	r.GET("/product/:id", pc.GetProduct)

	req, _ := http.NewRequest(http.MethodGet, "/product/100500", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Продукт не найден")
}

func TestProductController_GetProduct_Error(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)

	mockRepo.On("GetProduct", 1).Return(models.Product{}, assert.AnError)

	r := gin.Default()
	pc := NewProductController(nil, mockRepo, mockService)
	r.GET("/product/:id", pc.GetProduct)

	req, _ := http.NewRequest(http.MethodGet, "/product/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Не удалось получить данные о товаре")
}

func TestProductController_GetProductsByCategory_Success(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)
	limit := 10

	products := []api.ProductsSectionResponse{
		{ID: 1, Title: "Product 1", Price: 100, Photo: "/item1.jpg"},
		{ID: 2, Title: "Product 2", Price: 200, Photo: "/item2.jpg"},
	}

	mockService.On("GetLimitParam", mock.Anything).Return(limit, nil)
	mockRepo.On("GetProductsByCategory", 1, limit).Return(products, nil)

	r := gin.Default()
	pc := NewProductController(nil, mockRepo, mockService)
	r.GET("/categories/:id/products", pc.GetProductsByCategory)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/categories/1/products?limit=%d", limit), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product 1")
}

func TestProductController_GetProductsByCategory_InvalidCategoryID(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)
	mockService.On("GetLimitParam", mock.Anything).Return(10, nil)

	r := gin.Default()
	pc := NewProductController(nil, mockRepo, mockService)
	r.GET("/categories/:id/products", pc.GetProductsByCategory)

	req, _ := http.NewRequest(http.MethodGet, "/categories/abc/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Неверный формат ID категории")
}

func TestProductController_GetProductsByCategory_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)
	limit := 10

	mockService.On("GetLimitParam", mock.Anything).Return(limit, nil)
	mockRepo.On("GetProductsByCategory", 1, limit).Return([]api.ProductsSectionResponse{}, gorm.ErrRecordNotFound)

	r := gin.Default()
	pc := NewProductController(nil, mockRepo, mockService)
	r.GET("/categories/:id/products", pc.GetProductsByCategory)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/categories/1/products?limit=%d", limit), nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Категория не найдена")

	mockRepo.AssertExpectations(t)
	mockService.AssertExpectations(t)
}

func TestProductController_GetProductsByCategory_Error(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)
	limit := 10

	mockService.On("GetLimitParam", mock.Anything).Return(limit, nil)
	mockRepo.On("GetProductsByCategory", 1, limit).Return([]api.ProductsSectionResponse{}, assert.AnError)

	r := gin.Default()
	pc := NewProductController(nil, mockRepo, mockService)
	r.GET("/categories/:id/products", pc.GetProductsByCategory)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/categories/1/products?limit=%d", limit), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Ошибка при получении данных о товарах")
}
