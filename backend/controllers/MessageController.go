package controllers

import (
	"backend/infrastructure/rabbitmq"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"regexp"
)

type MessageController struct {
	rmq *rabbitmq.RabbitMQ
}

func NewMessageController(rmq *rabbitmq.RabbitMQ) *MessageController {
	return &MessageController{rmq}
}

func (mc *MessageController) SendLeadMessage(c *gin.Context) {
	var formData models.LeadForm

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	phoneRegex := `^\+?[0-9]{10,15}$`
	match, _ := regexp.MatchString(phoneRegex, formData.Phone)
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат номера телефона"})
		return
	}

	message := fmt.Sprintf("Новая заявка: %s, телефон: %s, комментарий: %s", formData.Name, formData.Phone, formData.Message)

	err := mc.rmq.PublishMessage(os.Getenv("RABBITMQ_LEAD_QUEUE_NAME"), message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка отправки данных в очередь: %s", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Форма успешно отправлена"})
}
