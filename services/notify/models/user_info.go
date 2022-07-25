package models

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Icon  string `json:"icon"`
}

type MessageInfo struct {
	UserID       string       `json:"userID"`
	Company      string       `json:"company"`
	Notification Notification `json:"notification"`
}
