package dto

import "time"

type AllChats struct {
	ChatId    int64     `json:"chat_id"`
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	LineText  string    `json:"line_text"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatsLastMsgs struct {
	LastMessage    int64     `json:"LastMessage"`
	ConversationID int64     `json:"conversation_id"`
	MessageID      int64     `json:"message_id"`
	FromUser       int64     `json:"from_user"`
	Text           string    `json:"text"`
	CreatedAt      time.Time `json:"created_at"`
}

type ChatMessage struct {
	MessageId      int64     `json:"message_id"`
	FromUser       int64     `json:"from_user"`
	Text           string    `json:"text"`
	CreatedAt      time.Time `json:"created_at"`
	ConversationId int64     `json:"conversation_id"`
}
