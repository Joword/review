package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserMessage struct {
	ID             int    `gorm:"primary_key" json:"id"`
	Username       string `json:"username"`
	Nickname       string `json:"nickname"`
	Email          string `json:"email"`
	Organization   string `json:"organization"`
	Activetime     string `json:"active_time"`
	ChatStatus     string `json:"chat_status"`
	MaxConvCount   int    `json:"max_conv_count"`
	IsSuperuser    int    `json:"is_superuser"`
	HashedPassword string `json:"hashed_password"`
}
