package handlers

import (
	"college-bot/buttons"
	"college-bot/database"
	"database/sql"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Основной обработчик обновлений(сообщение или кнопка)
func HandleUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel, db *sql.DB) {
	keyboard := buttons.CreatMainMenu()

	for update := range updates {
		if update.Message != nil {
			HandleMessage(bot, update.Message, keyboard)
		}
		if update.CallbackQuery != nil {
			HandleCallBackQuery(bot, update.CallbackQuery, db)
		}
	}
}

// обработчик сообщений и команд
func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, keyboard tgbotapi.InlineKeyboardMarkup) {
	if message.IsCommand() {
		switch message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Я бот колледжа. Чем могу помочь?")
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		case "help":
			fmt.Println("Если у вас возникла проблема при использовании бота, напишите сюда: ссылка на другого бота)")
		case "хуй":
			fmt.Println("Сам ты хуй")
		default:
			msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда. Используйте /help для получения помощи или /start для использования бота.")
			bot.Send(msg)
		}

	}
}

// Обработчик кнопок
func HandleCallBackQuery(bot *tgbotapi.BotAPI, callBackQuery *tgbotapi.CallbackQuery, db *sql.DB) {
	data := callBackQuery.Data
	switch data {
	case "about":
		msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, "До")
		bot.Send(msg)
	case "schedule":
		day := "понедельник" // временно
		week := "числитель"  // временно
		scheduleText, err := database.GetSchedule(db, day, week)
		if err != nil {
			msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, "Ошибка при инициализации расписания")
			bot.Send(msg)
			return
		}
		msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, scheduleText)
		bot.Send(msg)
	case "other":
		keyboard := buttons.OtherMenu()
		msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, "Дополнительные функции")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	case "cencel":
		keyboardMenu := buttons.CreatMainMenu()
		msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, "Главное меню")
		msg.ReplyMarkup = keyboardMenu
		bot.Send(msg)

	case "reminder":
		msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, "Напоминание")
		bot.Send(msg)
	case "conspects":

	default:
		msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, "Неизвестная команда. Используйте /help для получения помощи или /start для использования бота.")
		bot.Send(msg)
	}

	callback := tgbotapi.NewCallback(callBackQuery.ID, "")
	bot.Request(callback)
}
