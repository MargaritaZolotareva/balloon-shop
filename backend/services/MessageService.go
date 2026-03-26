package services

import (
	"backend/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageService interface {
	SendMessage(formData models.LeadForm) error
}

type MessageServiceImpl struct {
	tgBot *tgbotapi.BotAPI
}

func NewMessageService(api *tgbotapi.BotAPI) MessageService {
	return &MessageServiceImpl{tgBot: api}
}

func (ns *MessageServiceImpl) SendMessage(formData models.LeadForm) error {
	message := fmt.Sprintf("Новая заявка: %s, телефон: %s, комментарий: %s", formData.Name, formData.Phone, formData.Message)

	log.Printf("Received message: %s", message)
	chatIDs := os.Getenv("TG_BOT_ADMIN_CHAT_IDS")
	for _, chatID := range strings.Split(chatIDs, ",") {
		chatID, err := strconv.ParseInt(chatID, 10, 64)
		tgMessage := tgbotapi.NewMessage(chatID, message)
		_, err = ns.tgBot.Send(tgMessage)
		if err != nil {
			log.Printf("Failed to send message to Telegram: %v", err)
		} else {
			log.Printf("Message sent to Telegram: %s", message)
		}

		return err
	}

	return nil
}
