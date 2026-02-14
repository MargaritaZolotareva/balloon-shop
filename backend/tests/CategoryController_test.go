package tests

import (
	"backend/controllers"
	"backend/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockCategoryRepo struct {
	mock.Mock
}

func (m *MockCategoryRepo) GetCategory(id int) (models.Category, error) {
	args := m.Called(id)
	return args.Get(0).(models.Category), args.Error(1)
}

func TestCategoryController_GetCategory_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepo)
	category := models.Category{ID: 1, Title: "Test Category"}
	mockRepo.On("GetCategory", 1).Return(category, nil)

	r := gin.Default()
	cc := controllers.NewCategoryController(nil, mockRepo)
	r.GET("/categories/:id", cc.GetCategory)

	req, _ := http.NewRequest(http.MethodGet, "/category/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Category")

	mockRepo.AssertExpectations(t)
}

func TestCategoryController_GetCategory_InvalidID(t *testing.T) {
	mockRepo := new(MockCategoryRepo)

	r := gin.Default()
	cc := controllers.NewCategoryController(nil, mockRepo)
	r.GET("/categories/:id", cc.GetCategory)

	req, _ := http.NewRequest(http.MethodGet, "/category/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Неверный формат ID категории")

	mockRepo.AssertExpectations(t)
}

func TestCategoryController_GetCategory_NotFound(t *testing.T) {
	mockRepo := new(MockCategoryRepo)
	mockRepo.On("GetCategory", 100500).Return(models.Category{}, gorm.ErrRecordNotFound)

	r := gin.Default()
	cc := controllers.NewCategoryController(nil, mockRepo)
	r.GET("/categories/:id", cc.GetCategory)

	req, _ := http.NewRequest(http.MethodGet, "/category/100500", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Товар не найден")

	mockRepo.AssertExpectations(t)
}

func TestCategoryController_GetCategory_DBError(t *testing.T) {
	mockRepo := new(MockCategoryRepo)

	mockRepo.On("GetCategory", 1).Return(models.Category{}, errors.New("Database Error"))

	r := gin.Default()
	cc := controllers.NewCategoryController(nil, mockRepo)
	r.GET("/categories/:id", cc.GetCategory)

	req, _ := http.NewRequest(http.MethodGet, "/category/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Не удалось получить данные о категории")
}
