package models

// Task model struct
type Task struct {
	ID       int64  `json:"id" db:"id"`
	Text     string `json:"text" db:"text"`
	Day      string `json:"day" db:"day"`
	Reminder bool   `json:"reminder" db:"reminder"`
}
