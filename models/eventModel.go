package model

import (
	"time"

	"gorm.io/datatypes"
)

type Event struct {
	EventID    uint           `gorm:"primaryKey" json:"id_event"`
	UserID     int            `json:"user_id"`
	Status     int            `json:"status"`
	TitleEvent string         `json:"title_event"`
	DateEvent  datatypes.Date `json:"date_event"`
	StartEvent string         `json:"start_event"`
	EndEvent   string         `json:"end_event"`
	Location   string         `gorm:"size:50" json:"lokasi"`
	User       User           `gorm:"foreignKey:UserID;references:UserID"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}
