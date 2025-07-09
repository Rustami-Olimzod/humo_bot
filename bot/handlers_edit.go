package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"humo_bot/db"
	"strconv"
	"strings"
	"time"
)

func handleEditInit(bot *tgbotapi.BotAPI, update tgbotapi.Update, dbUser db.User, state *UserLateState) {
	state.CurrentAction = "editing_select"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите заявку для изменения (введите номер):\n"+getUserEventsList(dbUser))
	msg.ReplyMarkup = getBackKeyboard()
	bot.Send(msg)
}

func handleEditSelect(bot *tgbotapi.BotAPI, update tgbotapi.Update, dbUser db.User, state *UserLateState) {
	selectedNum, err := strconv.Atoi(update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if err != nil {
		msg.Text = "Неверный номер заявки. Пожалуйста, введите число из списка."
		bot.Send(msg)
		return
	}
	var events []db.Event
	db.DB.Where("user_id = ?", dbUser.ID).Order("created_at desc").Limit(10).Find(&events)
	if selectedNum < 1 || selectedNum > len(events) {
		msg.Text = "Заявка с таким номером не найдена. Пожалуйста, выберите другой номер."
		bot.Send(msg)
		return
	}
	event := events[selectedNum-1]
	state.EditingEventID = event.ID
	state.CurrentAction = "editing_field"
	msg.Text = fmt.Sprintf("Выберите поле для изменения заявки #%d:\nТип: %s\nДата: %s\nКомментарий: %s\nМинуты: %v",
		selectedNum, event.EventType, event.DateFrom.Format("02.01.2006"), event.Comment,
		func() interface{} {
			if event.Minutes != nil {
				return *event.Minutes
			}
			return "нет"
		}())
	msg.ReplyMarkup = getEditFieldsKeyboard()
	bot.Send(msg)
}

func handleEditField(bot *tgbotapi.BotAPI, update tgbotapi.Update, dbUser db.User, state *UserLateState) {
	switch strings.ToLower(update.Message.Text) {
	case "назад":
		handleEditInit(bot, update, dbUser, state)
	case "причину опоздания":
		state.CurrentAction = "editing_type"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите новый тип заявки:")
		msg.ReplyMarkup = getTypeKeyboard()
		bot.Send(msg)
	case "дата":
		state.CurrentAction = "editing_date"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите новую дату (ДД.ММ.ГГГГ):")
		msg.ReplyMarkup = getDateKeyboard()
		bot.Send(msg)
	case "комментарий":
		state.CurrentAction = "editing_comment"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите новый комментарий:")
		msg.ReplyMarkup = getBackKeyboard()
		bot.Send(msg)
	case "время":
		state.CurrentAction = "editing_minutes"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите новое количество минут:")
		msg.ReplyMarkup = GetTimeKeyboard()
		bot.Send(msg)
	}
}

func handleEditDate(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	text := update.Message.Text
	lowerText := strings.ToLower(text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	var date time.Time
	var err error
	switch lowerText {
	case "сегодня":
		date = time.Now()
	case "завтра":
		date = time.Now().AddDate(0, 0, 1)
	default:
		date, err = time.Parse("02.01.2006", text)
		if err != nil {
			msg.Text = "Неверный формат даты. Пожалуйста, укажите дату в формате ДД.ММ.ГГГГ"
			bot.Send(msg)
			return
		}
		today := time.Now().Truncate(24 * time.Hour)
		inputDay := date.Truncate(24 * time.Hour)
		if inputDay.Before(today) {
			msg.Text = "Нельзя выбирать прошедшую дату! Можно только сегодня или позже."
			bot.Send(msg)
			return
		}
	}
	var event db.Event
	db.DB.First(&event, state.EditingEventID)
	event.DateFrom = date
	event.DateTo = date
	db.DB.Save(&event)
	msg.Text = "Дата заявки успешно изменена."
	msg.ReplyMarkup = GetMainKeyboard()
	state.CurrentAction = ""
	bot.Send(msg)
}

func handleEditType(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	if strings.ToLower(update.Message.Text) == "назад" {
		state.CurrentAction = "editing_field"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите поле для изменения заявки:")
		msg.ReplyMarkup = getEditFieldsKeyboard()
		bot.Send(msg)
		return
	}
	var event db.Event
	db.DB.First(&event, state.EditingEventID)
	event.EventType = update.Message.Text
	db.DB.Save(&event)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Тип заявки успешно изменен.")
	msg.ReplyMarkup = GetMainKeyboard()
	state.CurrentAction = ""
	bot.Send(msg)
}

func handleEditComment(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	if strings.ToLower(update.Message.Text) == "назад" {
		state.CurrentAction = "editing_field"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите поле для изменения заявки:")
		msg.ReplyMarkup = getEditFieldsKeyboard()
		bot.Send(msg)
		return
	}
	var event db.Event
	db.DB.First(&event, state.EditingEventID)
	event.Comment = update.Message.Text
	db.DB.Save(&event)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Комментарий заявки успешно изменен.")
	msg.ReplyMarkup = GetMainKeyboard()
	state.CurrentAction = ""
	bot.Send(msg)
}

func handleEditMinutes(bot *tgbotapi.BotAPI, update tgbotapi.Update, state *UserLateState) {
	if strings.ToLower(update.Message.Text) == "назад" {
		state.CurrentAction = "editing_field"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите поле для изменения заявки:")
		msg.ReplyMarkup = getEditFieldsKeyboard()
		bot.Send(msg)
		return
	}
	minutes := parseTime(update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if minutes == 0 {
		msg.Text = "Пожалуйста, выберите время из предложенных вариантов"
		msg.ReplyMarkup = GetTimeKeyboard()
		bot.Send(msg)
		return
	}
	var event db.Event
	db.DB.First(&event, state.EditingEventID)
	event.Minutes = &minutes
	db.DB.Save(&event)
	msg.Text = "Время опоздания успешно изменено."
	msg.ReplyMarkup = GetMainKeyboard()
	state.CurrentAction = ""
	bot.Send(msg)
}
