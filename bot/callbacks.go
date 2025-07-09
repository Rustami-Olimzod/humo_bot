package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"humo_bot/db"
	"log"
)

func handleCallbackQuery(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	_, _ = bot.Request(callback)

	chatID := update.CallbackQuery.Message.Chat.ID
	userID := update.CallbackQuery.From.ID

	switch update.CallbackQuery.Data {
	case "confirm":
		state, exists := userLateStates[userID]
		if !exists || state.PendingEvent == nil {
			bot.Send(tgbotapi.NewMessage(chatID, "❌ Не удалось подтвердить заявку: данные устарели. Пожалуйста, создайте заявку заново."))
			return
		}
		if err := db.DB.Create(state.PendingEvent).Error; err != nil {
			log.Println("Error creating event:", err)
			bot.Send(tgbotapi.NewMessage(chatID, "❌ Ошибка при сохранении заявки"))
			return
		}
		msg := tgbotapi.NewMessage(chatID, "✅ Заявка подтверждена и отправлена.")
		msg.ReplyMarkup = GetMainKeyboard()
		bot.Send(msg)
		delete(userLateStates, userID)
	case "cancel":
		keyboard := getCancelKeyboard()
		msg := tgbotapi.NewMessage(chatID, "❌ Заявка отменена.")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
		delete(userLateStates, userID)
	}
}
