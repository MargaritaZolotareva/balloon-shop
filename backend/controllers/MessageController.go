package controllers

import (
	"backend/api"
	"backend/infrastructure/metrics"
	"backend/models"
	"backend/services"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	ns services.MessageService
}

func NewMessageController(ns services.MessageService) *MessageController {
	return &MessageController{ns}
}

func (mc *MessageController) SendLeadMessage(c *gin.Context) {
	var formData models.LeadForm

	if err := c.ShouldBindJSON(&formData); err != nil {
		metrics.Error400Counter.WithLabelValues("400").Inc()
		api.SendError(c, http.StatusBadRequest, "Неверные данные")
		return
	}

	phoneRegex := `^\+?[0-9]{10,15}$`
	match, _ := regexp.MatchString(phoneRegex, formData.Phone)
	if !match {
		metrics.Error400Counter.WithLabelValues("400").Inc()
		api.SendError(c, http.StatusBadRequest, "Неверный формат номера телефона")
		return
	}

	err := mc.ns.SendMessage(formData)
	if err != nil {
		metrics.Error500Counter.WithLabelValues("500").Inc()
		api.SendError(c, http.StatusInternalServerError, fmt.Sprintf("Ошибка отправки данных в очередь: %s", err))
		return
	}

	metrics.SuccessCounter.WithLabelValues("200").Inc()
	c.JSON(http.StatusOK, gin.H{"message": "Форма успешно отправлена"})
}
