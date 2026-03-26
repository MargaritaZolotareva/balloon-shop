package main

import (
	"backend/infrastructure"
	db2 "backend/infrastructure/db"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
	db := db2.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sqlDB: %v", err)
	}
	defer sqlDB.Close()

	bot := AuthBot()
	r := infrastructure.SetupRouter(db, bot)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	<-stop

	log.Println("Shutting down the server...")

	log.Println("Reporter closed. Server shutdown complete.")
}
