package buttons

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// CreateMainMenu создает главное меню с кнопками
func CreatMainMenu() tgbotapi.InlineKeyboardMarkup {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Информация о боте", "about")},
		{tgbotapi.NewInlineKeyboardButtonData("Расписание", "schedule")},
		{tgbotapi.NewInlineKeyboardButtonData("Доп функции", "other")},
	}

	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func OtherMenu() tgbotapi.InlineKeyboardMarkup {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Напоминание", "reminder")},
		{tgbotapi.NewInlineKeyboardButtonData("Конспекты", "")},
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func CancelMenu() tgbotapi.InlineKeyboardMarkup {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Назад", "cancel")},
		{tgbotapi.NewInlineKeyboardButtonData("Главное меню", "mainmenu")},
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}
