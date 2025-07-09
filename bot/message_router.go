package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"humo_bot/db"
	"strings"
)

func dispatchMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	user := update.Message.From

	// Сохраняем пользователя, если нет
	var dbUser db.User
	if err := db.DB.Where("telegram_id = ?", user.ID).First(&dbUser).Error; err != nil {
		dbUser = db.User{
			TelegramID: user.ID,
			Username:   user.UserName,
			FullName:   user.FirstName + " " + user.LastName,
		}
		db.DB.Create(&dbUser)
	}

	text := update.Message.Text
	lowerText := strings.ToLower(text)

	// Проверяем состояние пользователя
	state := getUserState(user.ID)

	// Обработка "назад"
	if lowerText == "назад" {
		handleBackCommand(bot, update, state)
		return
	}

	switch {
	case lowerText == "/start":
		handleStart(bot, update)

	case lowerText == "⏰ опоздание":
		handleLate(bot, update, state)

	case lowerText == "✏️ изменить":
		handleEditInit(bot, update, dbUser, state)

	case state.CurrentAction == "editing_select":
		handleEditSelect(bot, update, dbUser, state)

	case state.CurrentAction == "editing_field":
		handleEditField(bot, update, dbUser, state)

	case state.CurrentAction == "editing_date":
		handleEditDate(bot, update, state)

	case state.CurrentAction == "editing_type":
		handleEditType(bot, update, state)

	case state.CurrentAction == "editing_comment":
		handleEditComment(bot, update, state)

	case state.CurrentAction == "editing_minutes":
		handleEditMinutes(bot, update, state)

	case state.CurrentAction == "late" && !state.IsManualTime && (strings.Contains(lowerText, "минут") || strings.Contains(lowerText, "час")):
		handleLateTime(bot, update, state)

	case state.CurrentAction == "late" && lowerText == "указать вручную":
		handleLateManualRequest(bot, update, state)

	case state.CurrentAction == "late_manual_time":
		handleLateManualTime(bot, update, state)

	case state.CurrentAction == "late_reason":
		handleLateReason(bot, update, state)

	case state.CurrentAction == "late_comment":
		handleLateComment(bot, update, dbUser, state)

	case lowerText == "📋 история":
		showHistory(bot, update, dbUser)

	case lowerText == "удалить сегодняшние заявки":
		deleteAllToday(bot, update, dbUser)

	default:
		msg.Text = "Пожалуйста, выберите действие с клавиатуры."
		msg.ReplyMarkup = GetMainKeyboard()
		bot.Send(msg)
	}
}
