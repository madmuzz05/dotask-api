package model

import (
	"time"

	"gorm.io/datatypes"
)

type Task struct {
	TaskID    uint           `gorm:"primaryKey" json:"id_task"`
	TitleTask string         `json:"title_tsk"`
	DateTask  datatypes.Date `json:"date_task"`
	StartTask string         `json:"start_task"`
	EndTask   string         `json:"end_task"`
	Category  string         `gorm:"size;50"json:"category"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
