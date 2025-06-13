package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleBackCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню:")
	keyboard := GetMainKeyboard()
	keyboard.ResizeKeyboard = true
	msg.ReplyMarkup = keyboard
	state.CurrentAction = ""
	bot.Send(msg)
}
