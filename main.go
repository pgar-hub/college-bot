package main

import (
	"college-bot/database"
	"college-bot/handlers"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	//загрузка переменных окружения из .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	//подключение бд
	connStr := os.Getenv("connStr")
	if connStr == "" {
		panic("connStr environment variable is not set")
	}

	db, err := database.ConnectDB(connStr)
	if err != nil {
		log.Fatal("connect:", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("Ошибка ping БД:", err)
	}
	log.Println("Подключение к БД успешно установлено")

	// ✅ ВЫЗОВИТЕ SeedSchedule ПЕРЕД запуском бота
	if err := database.CreateTable(db); err != nil {
		log.Fatal("create table:", err)
	}
	if err := database.SeedSchedule(db); err != nil {
		log.Fatal("seed schedule:", err)
	}

	//токен бота
	token := os.Getenv("TOKEN")

	if token == "" {
		panic("API_TOKEN environment variable is not set")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	handlers.HandleUpdates(bot, updates, db)
}
