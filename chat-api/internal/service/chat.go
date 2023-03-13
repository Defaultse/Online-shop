package service

import (
	"chat-project-go/internal/datastruct"
	"chat-project-go/internal/repository"
)

type ChatService interface {
	GetAllChatsById(userId string) ([]datastruct.ChatsLastMsgs, error)
	StartNewConversation(fromUser string, toUser string, text string) error
	SendMessageToConversation(userId string, conversation_id string, message string) (int64, error)
	GetChatMessages(userId string, conversationId string) ([]datastruct.ChatMessage, error)
	GetConversationMembersIDs(conversationID string) ([]string, error)
}

type chatService struct {
	chatRepo repository.ChatRepositoryContract
}

func NewChatService(chatRepo repository.ChatRepositoryContract) *chatService {
	return &chatService{
		chatRepo: chatRepo,
	}
}

func (c chatService) StartNewConversation(fromUser string, toUser string, text string) error {
	conversationId, err := c.chatRepo.CreateNewConverasation(fromUser, toUser)

	if err != nil {
		return err
	}

	_, err = c.chatRepo.CreateMessage(fromUser, conversationId, text)

	if err != nil {
		return err
	}

	return nil
}

func (c chatService) SendMessageToConversation(userId string, conversationId string, message string) (int64, error) {
	var msgId int64

	userIsInConversation, err := c.chatRepo.CheckUserInConversation(userId, conversationId)

	if err != nil {
		return msgId, err
	}

	if userIsInConversation {
		msgId, err := c.chatRepo.CreateMessage(userId, conversationId, message)

		if err != nil {
			return msgId, err
		}
	}

	return msgId, nil
}

func (c chatService) GetConversationMembersIDs(conversationID string) ([]string, error) {
	membersIDs, err := c.chatRepo.GetConversationMemebersIDs(conversationID)

	if err != nil {
		return nil, err
	}

	return membersIDs, nil
}

func (c chatService) GetAllChatsById(userId string) ([]datastruct.ChatsLastMsgs, error) {
	return c.chatRepo.GetAllChatsByUserId(userId)
}

func (c chatService) GetChatMessages(userId string, conversationId string) ([]datastruct.ChatMessage, error) {
	var chatMsgs []datastruct.ChatMessage

	userIsInConversation, err := c.chatRepo.CheckUserInConversation(userId, conversationId)

	if err != nil {
		return nil, err
	}

	if userIsInConversation {
		chatMsgs, err = c.chatRepo.GetConversationMessages(conversationId)

		if err != nil {
			return nil, err
		}

		return chatMsgs, nil
	}

	return nil, nil
}
