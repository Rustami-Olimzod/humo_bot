package bot

import "humo_bot/db"

type UserLateState struct {
	CurrentAction  string
	TempMinutes    int
	TempReason     string
	TempComment    string
	IsManualTime   bool
	PendingEvent   *db.Event
	EditingEventID uint
}

var userLateStates = make(map[int64]*UserLateState)

func getUserState(userID int64) *UserLateState {
	state, exists := userLateStates[userID]
	if !exists {
		state = &UserLateState{}
		userLateStates[userID] = state
	}
	return state
}
