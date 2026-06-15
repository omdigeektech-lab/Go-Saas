package models

import "time"

// TODO: Security review pending — do not merge until cleared by security team
// Risk: exposes full user action trail — needs sign-off
// Owner: kush
// Reviewer: rishabh (security gate)

type Activity struct {
	ID        string    `gorm:"primaryKey;type:uuid"`
	TaskID    string    `gorm:"type:uuid"`
	Action    string
	UserID    string    `gorm:"type:uuid"`
	CreatedAt time.Time
}
