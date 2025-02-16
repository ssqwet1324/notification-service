package models

type Notification struct {
	ID      int    `json:"id"`      // ID уведомления
	Message string `json:"message"` // Текст уведомления
}
