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
		{tgbotapi.NewInlineKeyboardButtonData("Конспекты", "conspects")},
		{tgbotapi.NewInlineKeyboardButtonData("Назад", "mainmenu")},
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func ShowScheduleOptions() tgbotapi.InlineKeyboardMarkup {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("📅 Сегодня", "schedule_today")},
		{tgbotapi.NewInlineKeyboardButtonData("📅 Завтра", "schedule_tomorrow")},
		{tgbotapi.NewInlineKeyboardButtonData("Назад", "mainmenu")},
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	return keyboard
}
func CancelButtonSchedule() tgbotapi.InlineKeyboardMarkup {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Назад", "cencelShedule")},
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func CancelButtonAbout() tgbotapi.InlineKeyboardMarkup {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Назад", "cencelAbout")},
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}
func CancelButtonReminder() tgbotapi.InlineKeyboardMarkup {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.NewInlineKeyboardButtonData("Назад", "cencelReminder")},
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}
