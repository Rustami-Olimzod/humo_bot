package bot

import (
	"fmt"
	"humo_bot/db"
	"strconv"
	"strings"
)

func getUserEventsList(user db.User) string {
	var events []db.Event
	db.DB.Where("user_id = ?", user.ID).Order("created_at desc").Limit(10).Find(&events)
	var sb strings.Builder
	for i, e := range events {
		sb.WriteString(fmt.Sprintf("#%d - %s (%s)\n", i+1, e.EventType, e.CreatedAt.Format("02.01.2006")))
	}
	return sb.String()
}

func parseTime(text string) int {
	switch strings.ToLower(text) {
	case "5 минут":
		return 5
	case "10 минут":
		return 10
	case "15 минут":
		return 15
	case "30 минут":
		return 30
	case "1 час":
		return 60
	case "2 часа":
		return 120
	default:
		return 0
	}
}

func parseManualTime(text string) (int, error) {
	text = strings.ToLower(strings.TrimSpace(text))
	var cleaned strings.Builder
	for _, r := range text {
		if (r >= '0' && r <= '9') || r == '.' {
			cleaned.WriteRune(r)
		}
	}
	value, err := strconv.ParseFloat(cleaned.String(), 64)
	if err != nil {
		return 0, fmt.Errorf("не удалось распознать число")
	}
	if strings.Contains(text, "час") {
		return int(value * 60), nil
	}
	return int(value), nil
}
