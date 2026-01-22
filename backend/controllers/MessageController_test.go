package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockRabbitMQ struct {
	mock.Mock
}

func (m *MockRabbitMQ) PublishMessage(queueName string, message string) error {
	args := m.Called(queueName, message)
	return args.Error(0)
}

func TestMessageController_SendLeadMessage_Success(t *testing.T) {
	mockRMQ := new(MockRabbitMQ)
	mockRMQ.On("PublishMessage", mock.Anything, mock.Anything).Return(nil)

	r := gin.Default()
	mc := NewMessageController(mockRMQ)
	r.POST("/lead-form", mc.SendLeadMessage)

	req, _ := http.NewRequest(http.MethodPost, "/lead-form", strings.NewReader(`{"name": "John", "phone": "+1234567890", "message": "Hello"}`))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Форма успешно отправлена")
}

func TestMessageController_SendLeadMessage_InvalidJSON(t *testing.T) {
	mockRMQ := new(MockRabbitMQ)

	r := gin.Default()
	mc := NewMessageController(mockRMQ)
	r.POST("/lead-form", mc.SendLeadMessage)

	req, _ := http.NewRequest(http.MethodPost, "/lead-form", strings.NewReader(`{"name": "John"}`))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Неверные данные")
}

func TestMessageController_SendLeadMessage_InvalidPhoneFormat(t *testing.T) {
	mockRMQ := new(MockRabbitMQ)

	r := gin.Default()
	mc := NewMessageController(mockRMQ)
	r.POST("/lead-form", mc.SendLeadMessage)

	req, _ := http.NewRequest(http.MethodPost, "/lead-form", strings.NewReader(`{"name": "John", "phone": "invalidPhone", "message": "Hello"}`))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Неверный формат номера телефона")
}

func TestMessageController_SendLeadMessage_RabbitMQError(t *testing.T) {
	mockRMQ := new(MockRabbitMQ)

	mockRMQ.On("PublishMessage", mock.Anything, mock.Anything).Return(fmt.Errorf("ошибка RabbitMQ"))

	r := gin.Default()
	mc := NewMessageController(mockRMQ)
	r.POST("/lead-form", mc.SendLeadMessage)

	req, _ := http.NewRequest(http.MethodPost, "/lead-form", strings.NewReader(`{"name": "John", "phone": "+1234567890", "message": "Hello"}`))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Ошибка отправки данных в очередь")
}
