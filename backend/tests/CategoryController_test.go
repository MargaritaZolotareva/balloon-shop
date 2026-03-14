package tests

import (
	"backend/api"
	"backend/controllers"
	"backend/models"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockCategoryRepo struct {
	mock.Mock
}

func (m *MockCategoryRepo) GetCategoryBySlug(categorySlug string) (models.Category, error) {
	args := m.Called(categorySlug)
	return args.Get(0).(models.Category), args.Error(1)
}

func (m *MockCategoryRepo) GetAllCategories() ([]api.CategoriesResponse, error) {
	args := m.Called()
	return args.Get(0).([]api.CategoriesResponse), args.Error(1)
}

func (m *MockCategoryRepo) GetHomepageCategories() (api.HomepageResponse, error) {
	args := m.Called()
	return args.Get(0).(api.HomepageResponse), args.Error(1)
}

func TestCategoryController_GetCategory_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepo)
	category := models.Category{ID: 1, Title: "Test Category", PhotoID: 1, Slug: "1"}
	mockRepo.On("GetCategoryBySlug", category.Slug).Return(category, nil)

	r := gin.Default()
	cc := controllers.NewCategoryController(nil, mockRepo)
	r.GET("/categories/:slug", cc.GetCategoryBySlug)

	req, _ := http.NewRequest(http.MethodGet, "/categories/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Category")

	mockRepo.AssertExpectations(t)
}

func TestCategoryController_GetCategory_StringSlug(t *testing.T) {
	mockRepo := new(MockCategoryRepo)
	category := models.Category{ID: 1, Title: "Test Category", PhotoID: 1, Slug: "abc"}
	mockRepo.On("GetCategoryBySlug", category.Slug).Return(category, nil)

	r := gin.Default()
	cc := controllers.NewCategoryController(nil, mockRepo)
	r.GET("/categories/:slug", cc.GetCategoryBySlug)

	req, _ := http.NewRequest(http.MethodGet, "/categories/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Category")

	mockRepo.AssertExpectations(t)
}

func TestCategoryController_GetCategory_NotFound(t *testing.T) {
	mockRepo := new(MockCategoryRepo)
	mockRepo.On("GetCategoryBySlug", "100500").Return(models.Category{}, gorm.ErrRecordNotFound)

	r := gin.Default()
	cc := controllers.NewCategoryController(nil, mockRepo)
	r.GET("/categories/:slug", cc.GetCategoryBySlug)

	req, _ := http.NewRequest(http.MethodGet, "/categories/100500", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Категория не найдена")

	mockRepo.AssertExpectations(t)
}

func TestCategoryController_GetCategory_DBError(t *testing.T) {
	mockRepo := new(MockCategoryRepo)

	mockRepo.On("GetCategoryBySlug", "1").Return(models.Category{}, errors.New("Database Error"))

	r := gin.Default()
	cc := controllers.NewCategoryController(nil, mockRepo)
	r.GET("/categories/:slug", cc.GetCategoryBySlug)

	req, _ := http.NewRequest(http.MethodGet, "/categories/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Не удалось получить данные о категории")
}
