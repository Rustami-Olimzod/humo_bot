package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "👋 Добро пожаловать! Пожалуйста, выберите действие из списка.")
	msg.ReplyMarkup = GetMainKeyboard()
	bot.Send(msg)
}
