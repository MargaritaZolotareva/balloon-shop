package main

import (
	"backend/infrastructure"
	"backend/infrastructure/rabbitmq"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка при соединении с БД: %v", err)
	}

	return db
}

func InitRabbitMq() *rabbitmq.RabbitMQ {
	rmq, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	_, err = rmq.DeclareQueue(os.Getenv("RABBITMQ_LEAD_QUEUE_NAME"))
	if err != nil {
		log.Fatalf("failed to create queue: %v", err)
	}

	return rmq
}

func startConsumer(r *rabbitmq.RabbitMQ, bot *tgbotapi.BotAPI) {
	go func() {
		err := r.StartConsuming(os.Getenv("RABBITMQ_LEAD_QUEUE_NAME"), func(msg amqp.Delivery) {
			message := string(msg.Body)
			log.Printf("Received message: %s", message)
			chatIDStr := os.Getenv("TG_BOT_CHAT_ID")
			chatID, err := strconv.ParseInt(chatIDStr, 10, 64)

			tgMessage := tgbotapi.NewMessage(chatID, message)
			_, err = bot.Send(tgMessage)
			if err != nil {
				log.Printf("Failed to send message to Telegram: %v", err)
			} else {
				log.Printf("Message sent to Telegram: %s", message)
			}
		})
		if err != nil {
			log.Printf("Failed to start consuming: %v", err)
		}
	}()
}

func AuthBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_API_TOKEN"))
	if err != nil {
		log.Fatalf("Failed to create Telegram bot: %v", err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func main() {
	log.SetOutput(os.Stdout)
	db := InitDB()
	rmq := InitRabbitMq()
	defer rmq.Close()
	sqlDB, err := db.DB()
	if err != nil {
		sqlDB.Close()
	}
	r := infrastructure.SetupRouter(db, rmq)
	bot := AuthBot()

	startConsumer(rmq, bot)

	r.Run(":8080")
}
