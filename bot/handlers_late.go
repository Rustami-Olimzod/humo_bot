package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"humo_bot/db"
	"strings"
	"time"
)

func handleLate(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	state.CurrentAction = "late"
	state.IsManualTime = false
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "⏱ Укажите время опоздания:")
	msg.ReplyMarkup = GetTimeKeyboard()
	bot.Send(msg)
}

func handleLateTime(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	minutes := parseTime(update.Message.Text)
	if minutes == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, выберите время из предложенных вариантов")
		msg.ReplyMarkup = GetTimeKeyboard()
		bot.Send(msg)
		return
	}
	state.TempMinutes = minutes
	state.CurrentAction = "late_reason"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "📝 Пожалуйста, выберите причину опоздания:")
	msg.ReplyMarkup = GetLateReasonsKeyboard()
	bot.Send(msg)
}

func handleLateManualRequest(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	state.IsManualTime = true
	state.CurrentAction = "late_manual_time"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "⏱ Введите время опоздания (например: 45 мин или 1.5 час):")
	msg.ReplyMarkup = getBackKeyboard()
	bot.Send(msg)
}

func handleLateManualTime(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	minutes, err := parseManualTime(update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if err != nil {
		msg.Text = "Неверный формат. Примеры:\n- 45 (минуты)\n- 1.5 час\n- 2 часа"
	} else if minutes <= 0 {
		msg.Text = "Время должно быть больше 0"
	} else if minutes > 8*60 {
		msg.Text = "Слишком большое время опоздания"
	} else {
		state.TempMinutes = minutes
		state.IsManualTime = false
		state.CurrentAction = "late_reason"
		msg.Text = fmt.Sprintf("⏱ Время опоздания: %d минут. Пожалуйста, выберите причину опоздания!:", minutes)
		msg.ReplyMarkup = GetLateReasonsKeyboard()
	}
	bot.Send(msg)
}

func handleLateReason(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	state.TempReason = update.Message.Text
	state.CurrentAction = "late_comment"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "💬 Добавьте комментарий (или нажмите 'Пропустить'):")
	msg.ReplyMarkup = getSkipBackKeyboard()
	bot.Send(msg)
}

func handleLateComment(bot *tgbotapi.BotAPI, update tgbotapi.Update, dbUser db.User, state *UserLateState) {
	if strings.ToLower(update.Message.Text) == "пропустить" {
		state.TempComment = ""
	} else {
		state.TempComment = update.Message.Text
	}
	state.PendingEvent = &db.Event{
		UserID:    dbUser.ID,
		EventType: "Опоздание",
		Comment:   fmt.Sprintf("%s: %s", state.TempReason, state.TempComment),
		Minutes:   &state.TempMinutes,
		DateFrom:  time.Now(),
		DateTo:    time.Now(),
		Timestamp: time.Now().Unix(),
	}
	confirmationText := fmt.Sprintf("Подтвердите заявку на опоздание:\n\n⏱ Время: %d минут\n📝 Причина: %s\n💬 Комментарий: %s",
		state.TempMinutes, state.TempReason, state.TempComment)
	confirmMsg := tgbotapi.NewMessage(update.Message.Chat.ID, confirmationText)
	confirmMsg.ReplyMarkup = GetConfirmationInlineKeyboard()
	bot.Send(confirmMsg)
}
