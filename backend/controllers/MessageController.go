package controllers

import (
	"backend/api"
	"backend/infrastructure/metrics"
	"backend/infrastructure/rabbitmq"
	"backend/models"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"regexp"
)

type MessageController struct {
	rmq rabbitmq.MqInterface
}

func NewMessageController(rmq rabbitmq.MqInterface) *MessageController {
	return &MessageController{rmq}
}

func (mc *MessageController) SendLeadMessage(c *gin.Context) {
	var formData models.LeadForm

	if err := c.ShouldBindJSON(&formData); err != nil {
		sentry.CaptureException(err)
		metrics.Error400Counter.WithLabelValues("400").Inc()
		api.SendError(c, http.StatusBadRequest, "Неверные данные")
		return
	}

	phoneRegex := `^\+?[0-9]{10,15}$`
	match, _ := regexp.MatchString(phoneRegex, formData.Phone)
	if !match {
		sentry.CaptureMessage("Неверный формат номера телефона")
		metrics.Error400Counter.WithLabelValues("400").Inc()
		api.SendError(c, http.StatusBadRequest, "Неверный формат номера телефона")
		return
	}

	message := fmt.Sprintf("Новая заявка: %s, телефон: %s, комментарий: %s", formData.Name, formData.Phone, formData.Message)

	err := mc.rmq.PublishMessage(os.Getenv("RABBITMQ_LEAD_QUEUE_NAME"), message)
	if err != nil {
		sentry.CaptureException(err)
		metrics.Error500Counter.WithLabelValues("500").Inc()
		api.SendError(c, http.StatusInternalServerError, fmt.Sprintf("Ошибка отправки данных в очередь: %s", err))
		return
	}

	metrics.SuccessCounter.WithLabelValues("200").Inc()
	c.JSON(http.StatusOK, gin.H{"message": "Форма успешно отправлена"})
}
