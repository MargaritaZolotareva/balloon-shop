package tests

import (
	"backend/api"
	"backend/controllers"
	"backend/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockProductRepo struct {
	mock.Mock
}
type MockProductService struct {
	mock.Mock
}

func (m *MockProductRepo) GetProductBySlug(productSlug string) (models.Product, error) {
	args := m.Called(productSlug)
	return args.Get(0).(models.Product), args.Error(1)
}
func (m *MockProductRepo) GetSimilarProducts(categoryID, productID int) ([]api.SimilarProductResponse, error) {
	args := m.Called(categoryID, productID)
	return args.Get(0).([]api.SimilarProductResponse), args.Error(1)
}
func (m *MockProductRepo) GetProductsByCategory(categorySlug string, limit int) (api.GetProductsByCategoryResponse, error) {
	args := m.Called(categorySlug, limit)
	return args.Get(0).(api.GetProductsByCategoryResponse), args.Error(1)
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
		Slug:        "1",
	}

	mockRepo.On("GetProductBySlug", product.Slug).Return(product, nil)
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
	pc := controllers.NewProductController(nil, mockRepo, mockService)
	r.GET("/product/:slug", pc.GetProductBySlug)

	req, _ := http.NewRequest(http.MethodGet, "/product/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product")

	mockRepo.AssertExpectations(t)
	mockService.AssertExpectations(t)
}

func TestProductController_GetProduct_StringSlug(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)

	product := models.Product{
		ID:          1,
		Title:       "Test Product",
		Description: "Test description",
		Price:       100,
		CategoryId:  1,
		Slug:        "abc",
	}

	mockRepo.On("GetProductBySlug", product.Slug).Return(product, nil)
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
	pc := controllers.NewProductController(nil, mockRepo, mockService)
	r.GET("/product/:slug", pc.GetProductBySlug)

	req, _ := http.NewRequest(http.MethodGet, "/product/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product")

	mockRepo.AssertExpectations(t)
	mockService.AssertExpectations(t)
}

func TestProductController_GetProduct_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)

	mockRepo.On("GetProductBySlug", "100500").Return(models.Product{}, gorm.ErrRecordNotFound)

	r := gin.Default()
	pc := controllers.NewProductController(nil, mockRepo, mockService)
	r.GET("/product/:slug", pc.GetProductBySlug)

	req, _ := http.NewRequest(http.MethodGet, "/product/100500", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Продукт не найден")
}

func TestProductController_GetProduct_Error(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)

	mockRepo.On("GetProductBySlug", "1").Return(models.Product{}, assert.AnError)

	r := gin.Default()
	pc := controllers.NewProductController(nil, mockRepo, mockService)
	r.GET("/product/:slug", pc.GetProductBySlug)

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
	productsByCat := api.GetProductsByCategoryResponse{
		Products:      products,
		CategoryTitle: "Category Title",
	}

	mockService.On("GetLimitParam", mock.Anything).Return(limit, nil)
	mockRepo.On("GetProductsByCategory", "1", limit).Return(productsByCat, nil)

	r := gin.Default()
	pc := controllers.NewProductController(nil, mockRepo, mockService)
	r.GET("/categories/:slug/products", pc.GetProductsByCategory)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/categories/1/products?limit=%d", limit), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product 1")
}

func TestProductController_GetProductsByCategory_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockService := new(MockProductService)
	limit := 10

	mockService.On("GetLimitParam", mock.Anything).Return(limit, nil)
	mockRepo.On("GetProductsByCategory", "1", limit).Return(api.GetProductsByCategoryResponse{}, gorm.ErrRecordNotFound)

	r := gin.Default()
	pc := controllers.NewProductController(nil, mockRepo, mockService)
	r.GET("/categories/:slug/products", pc.GetProductsByCategory)

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
	mockRepo.On("GetProductsByCategory", "1", limit).Return(api.GetProductsByCategoryResponse{}, assert.AnError)

	r := gin.Default()
	pc := controllers.NewProductController(nil, mockRepo, mockService)
	r.GET("/categories/:slug/products", pc.GetProductsByCategory)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/categories/1/products?limit=%d", limit), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Ошибка при получении данных о товарах")
}
