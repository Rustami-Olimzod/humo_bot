package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Главный обработчик Telegram-апдейтов
func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		handleCallbackQuery(bot, &update)
		return
	}
	if update.Message == nil {
		return
	}
	dispatchMessage(bot, update)
}
