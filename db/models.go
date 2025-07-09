package db

import (
	"time"
)

type User struct {
	ID         uint  `gorm:"primaryKey"`
	TelegramID int64 `gorm:"uniqueIndex"`
	Username   string
	FullName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Event struct {
	ID              uint `gorm:"primaryKey"`
	UserID          uint
	EventType       string
	Comment         string
	Minutes         *int
	DateFrom        time.Time
	DateTo          time.Time
	CreatedAt       time.Time
	User            User      `gorm:"foreignKey:UserID"`
	ExpectedArrival time.Time `gorm:"column:expected_arrival"`
}
