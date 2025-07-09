package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"humo_bot/db"
	"strings"
	"time"
)

func showHistory(bot *tgbotapi.BotAPI, update tgbotapi.Update, user db.User) {
	var events []db.Event
	db.DB.Where("user_id = ? AND event_type = ?", user.ID, "Опоздание").Order("created_at desc").Limit(10).Find(&events)

	if len(events) == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "История опозданий пуста.")
		msg.ReplyMarkup = GetHistoryKeyboard()
		bot.Send(msg)
		return
	}

	var sb strings.Builder
	sb.WriteString("🕒 Последние опоздания:\n")
	for _, e := range events {
		date := e.CreatedAt.Format("02.01.2006 15:04")
		minutes := ""
		if e.Minutes != nil {
			minutes = fmt.Sprintf(" (%d мин)", *e.Minutes)
		}
		sb.WriteString(fmt.Sprintf("%s%s — %s\n", date, minutes, e.Comment))
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, sb.String())
	msg.ReplyMarkup = GetHistoryKeyboard()
	bot.Send(msg)
}

func deleteAllToday(bot *tgbotapi.BotAPI, update tgbotapi.Update, user db.User) {
	today := time.Now().Format("2006-01-02")
	var todaysEvents []db.Event
	db.DB.Where("user_id = ? AND DATE(created_at) = ?", user.ID, today).Find(&todaysEvents)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if len(todaysEvents) == 0 {
		msg.Text = "У вас нет заявок за сегодня, чтобы удалить."
		msg.ReplyMarkup = GetHistoryKeyboard()
	} else {
		db.DB.Where("user_id = ? AND DATE(created_at) = ?", user.ID, today).Delete(&db.Event{})
		msg.Text = fmt.Sprintf("Удалено заявок за сегодня: %d", len(todaysEvents))
		msg.ReplyMarkup = GetHistoryKeyboard()
	}
	bot.Send(msg)
}
