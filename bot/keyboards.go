package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("⏰ Опоздание"),
			tgbotapi.NewKeyboardButton("✏️ Изменить"),
			tgbotapi.NewKeyboardButton("📋 История"),
		),
	)
}

func GetTimeKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("5 минут"),
			tgbotapi.NewKeyboardButton("10 минут"),
			tgbotapi.NewKeyboardButton("15 минут"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("30 минут"),
			tgbotapi.NewKeyboardButton("1 час"),
			tgbotapi.NewKeyboardButton("2 часа"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Указать вручную"),
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
}

func GetLateReasonsKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("🚗 Пробки")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("🏡 Семейные обстоятельства")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("🚌 Общественный транспорт")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("🌧️ Погодные условия")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("⁉️ Другое "),
			tgbotapi.NewKeyboardButton("↩️ Назад"),
		),
	)
}

func GetHistoryKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Удалить сегодняшние заявки"),
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
}

func GetConfirmationInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ Подтвердить", "confirm"),
			tgbotapi.NewInlineKeyboardButtonData("❌ Отменить", "cancel"),
		),
	)
}

func getBackKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
}

func getSkipBackKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Пропустить"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
}

func getEditFieldsKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Причину опоздания"),
			tgbotapi.NewKeyboardButton("Дата"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Комментарий"),
			tgbotapi.NewKeyboardButton("Время"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
}

func getTypeKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Опоздание"),
			tgbotapi.NewKeyboardButton("Не приду"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Болезнь"),
			tgbotapi.NewKeyboardButton("Отпуск"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Командировка"),
			tgbotapi.NewKeyboardButton("По работе"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Другое"),
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
}

func getDateKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Сегодня"),
			tgbotapi.NewKeyboardButton("Завтра"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
}

func getCancelKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Оформить заявку заново"),
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
}
