package services

import (
	"backend/infrastructure/rabbitmq"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strconv"
	"strings"
)

func StartConsumer(r *rabbitmq.RabbitMQ, bot *tgbotapi.BotAPI) {
	go func() {
		err := r.StartConsuming(os.Getenv("RABBITMQ_LEAD_QUEUE_NAME"), func(msg amqp.Delivery) {
			message := string(msg.Body)
			log.Printf("Received message: %s", message)
			chatIDs := os.Getenv("TG_BOT_ADMIN_CHAT_IDS")
			for _, chatID := range strings.Split(chatIDs, ",") {
				chatID, err := strconv.ParseInt(chatID, 10, 64)
				tgMessage := tgbotapi.NewMessage(chatID, message)
				_, err = bot.Send(tgMessage)
				if err != nil {
					log.Printf("Failed to send message to Telegram: %v", err)
				} else {
					log.Printf("Message sent to Telegram: %s", message)
				}
			}
			msg.Ack(false)
		})
		if err != nil {
			log.Printf("Failed to start consuming: %v", err)
		}
	}()
}
