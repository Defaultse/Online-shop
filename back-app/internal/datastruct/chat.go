package datastruct

import (
	"time"
)

type AllChats struct {
	ChatId    int64
	Id        int64
	UserId    int64
	LineText  string
	CreatedAt time.Time
}

type ChatsLastMsgs struct {
	LastMessage    int64     `json:"LastMessage"`
	ProfileImage   string    `json:"profile_image"`
	Username       string    `json:"username"`
	ConversationID int64     `json:"conversation_id"`
	MessageID      int64     `json:"message_id"`
	FromUser       int64     `json:"from_user"`
	Text           string    `json:"text"`
	CreatedAt      time.Time `json:"created_at"`
}

type ChatMessage struct {
	Username       string
	MessageId      int64
	FromUser       int64
	Text           string
	CreatedAt      time.Time
	ConversationId int64
}
