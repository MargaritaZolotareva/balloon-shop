package main

import (
	"backend/infrastructure"
	"backend/infrastructure/rabbitmq"
	"backend/services"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
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

//func InitSentry() {
//	dsn := fmt.Sprintf("https://%s@%s/%s",
//		os.Getenv("SENTRY_PUBLIC_KEY"),
//		os.Getenv("SENTRY_HOST"),
//		os.Getenv("SENTRY_PROJECT_ID"))
//	err := sentry.Init(sentry.ClientOptions{
//		Dsn:              dsn,
//		TracesSampleRate: 1.0,
//	})
//	if err != nil {
//		log.Fatalf("sentry.Init: %s", err)
//	}
//
//	defer sentry.Flush(2 * time.Second)
//}

func initZipkin() (*zipkin.Tracer, reporter.Reporter) {
	rprtr := http.NewReporter(os.Getenv("ZIPKIN_URL"))

	log.Printf("Created reporter")

	endPoint, err := zipkin.NewEndpoint("my-service", "backend:8080")
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	tracer, err := zipkin.NewTracer(rprtr, zipkin.WithLocalEndpoint(endPoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %v", err)
	}

	return tracer, rprtr
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
	rmq := rabbitmq.InitRabbitMq()
	defer rmq.Close()
	sqlDB, err := db.DB()
	if err != nil {
		sqlDB.Close()
	}
	tracer, rprtr := initZipkin()
	r := infrastructure.SetupRouter(db, rmq, tracer)
	bot := AuthBot()

	services.StartConsumer(rmq, bot)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	<-stop

	log.Println("Shutting down the server...")
	rprtr.Close()
	log.Println("Reporter closed. Server shutdown complete.")
}
