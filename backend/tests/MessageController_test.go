package tests

import (
	"backend/controllers"
	"backend/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMessageService struct {
	mock.Mock
}

func (m *MockMessageService) SendMessage(formData models.LeadForm) error {
	args := m.Called(formData)
	return args.Error(0)
}

func TestMessageController_SendLeadMessage_Success(t *testing.T) {
	r := gin.Default()

	mockService := new(MockMessageService)
	mockService.On("SendMessage", mock.Anything).Return(nil)

	mc := controllers.NewMessageController(mockService)
	r.POST("/lead-form", mc.SendLeadMessage)

	req, _ := http.NewRequest(http.MethodPost, "/lead-form", strings.NewReader(`{"name": "John", "phone": "+1234567890", "message": "Hello"}`))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Форма успешно отправлена")
}

func TestMessageController_SendLeadMessage_InvalidJSON(t *testing.T) {
	r := gin.Default()

	mockService := new(MockMessageService)
	mockService.On("SendMessage", mock.Anything).Return(nil)

	mc := controllers.NewMessageController(mockService)
	r.POST("/lead-form", mc.SendLeadMessage)

	req, _ := http.NewRequest(http.MethodPost, "/lead-form", strings.NewReader(`{"name": "John"}`))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Неверные данные")
}

func TestMessageController_SendLeadMessage_InvalidPhoneFormat(t *testing.T) {
	r := gin.Default()

	mockService := new(MockMessageService)
	mockService.On("SendMessage", mock.Anything).Return(nil)

	mc := controllers.NewMessageController(mockService)
	r.POST("/lead-form", mc.SendLeadMessage)

	req, _ := http.NewRequest(http.MethodPost, "/lead-form", strings.NewReader(`{"name": "John", "phone": "invalidPhone", "message": "Hello"}`))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Неверный формат номера телефона")
}

func TestMessageController_SendLeadMessage_RabbitMQError(t *testing.T) {
	r := gin.Default()

	mockService := new(MockMessageService)
	mockService.On("SendMessage", mock.Anything).Return(fmt.Errorf("Ошибка отправки в ТГ"))

	mc := controllers.NewMessageController(mockService)
	r.POST("/lead-form", mc.SendLeadMessage)

	req, _ := http.NewRequest(http.MethodPost, "/lead-form", strings.NewReader(`{"name": "John", "phone": "+1234567890", "message": "Hello"}`))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Ошибка отправки данных в очередь")
}
