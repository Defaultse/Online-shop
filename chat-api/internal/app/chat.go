package app

import (
	"chat-project-go/internal/dto"
	"chat-project-go/pkg/websocket"
	"encoding/json"
	"fmt"
	"strconv"
)

func (s *Services) StartConversationWithUser(message websocket.Message) {

}

func (s *Services) SendMessageToConversation(message websocket.Message) {
	senderUserId := message.UserID
	conversationId := message.Body["conversation_id"].(string)
	text := message.Body["text"].(string)

	// must get inserted rows
	msgID, err := s.chatService.SendMessageToConversation(senderUserId, conversationId, text)

	if err != nil {
		fmt.Println(err)
	}

	senderUserIdInt64, _ := strconv.ParseInt(senderUserId, 10, 64)
	conversationIdInt64, _ := strconv.ParseInt(conversationId, 10, 64)

	createdMessage := dto.ChatMessage{
		MessageId:      msgID,
		FromUser:       senderUserIdInt64,
		Text:           text,
		ConversationId: conversationIdInt64,
	}

	jsonBytes, _ := json.Marshal(createdMessage)

	response := dto.WebsocketMsg2{
		Type: 5,
		Body: string(jsonBytes),
	}

	members, err := s.chatService.GetConversationMembersIDs(conversationId)

	for _, v := range members {
		if s.wsPool.Clients[v] != nil {
			s.wsPool.Clients[v].WriteJSON(response)
		}
	}
}

func (s *Services) GetConversationMsgs(message websocket.Message) {
	userId := message.UserID
	conversationId := message.Body["conversation_id"].(string)

	result, err := s.chatService.GetChatMessages(userId, conversationId)

	if err != nil {
		fmt.Println(err)
	}

	jsonBytes, _ := json.Marshal(result)

	response := dto.WebsocketMsg2{
		Type: GetConversationMsgs,
		Body: string(jsonBytes),
	}

	s.wsPool.Clients[userId].WriteJSON(response)

}

func (s *Services) GetAllChatsLastMsg(message websocket.Message) {
	userId := message.UserID

	result, err := s.chatService.GetAllChatsById(userId)

	if err != nil {
		return
	}

	jsonBytes, _ := json.Marshal(result)

	response := dto.WebsocketMsg2{
		Type: GettAllChatsLastMsg,
		Body: string(jsonBytes),
	}

	s.wsPool.Clients[userId].WriteJSON(response)
}
