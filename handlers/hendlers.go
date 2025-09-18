package handlers

import (
	"college-bot/buttons"
	"college-bot/photo"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Основной обработчик обновлений(сообщение или кнопка)
func HandleUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	keyboard := buttons.CreatMainMenu()

	for update := range updates {
		if update.Message != nil {
			HandleMessage(bot, update.Message, keyboard)
		}
		if update.CallbackQuery != nil {
			HandleCallBackQuery(bot, update.CallbackQuery)
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
func HandleCallBackQuery(bot *tgbotapi.BotAPI, callBackQuery *tgbotapi.CallbackQuery) {
	data := callBackQuery.Data
	switch data {
	case "about":
		msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, "пока нихуя нету")
		bot.Send(msg)
	case "schedule":
		a := photo.GetPhoto()
		if a == nil {
			log.Println("Фото не найдено")
			return
		}
		msg := tgbotapi.NewPhoto(callBackQuery.Message.Chat.ID, a)
		bot.Send(msg)
	case "other":
		msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, "пока нихуя нету")
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(callBackQuery.Message.Chat.ID, "Неизвестная команда. Используйте /help для получения помощи или /start для использования бота.")
		bot.Send(msg)
	}
	callback := tgbotapi.NewCallback(callBackQuery.ID, "")
	bot.Request(callback)
}
