package model

type Task struct {
	TaskID    uint   `gorm:"primaryKey" json:"id_task"`
	TitleTask string `json:"title_task`
	DateTask  string `json:"date_task"`
	StartTask string `json:"start_task"`
	EndTask   string `json:"end_task"`
	Company   string `json:"company"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
