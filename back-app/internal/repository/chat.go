package repository

import (
	"chat-project-go/internal/datastruct"
	"database/sql"
	"fmt"
)

type ChatRepositoryContract interface {
	GetAllChatsByUserId(userId string) ([]datastruct.ChatsLastMsgs, error)
	CreateMessage(from_user string, conversation_id string, text string) (int64, error)
	CreateNewConverasation(firstUserID string, secondUserID string) (string, error)
	GetConversationMessages(conversationId string) ([]datastruct.ChatMessage, error)
	CheckUserInConversation(userId string, conversationId string) (bool, error)
	GetConversationMemebersIDs(conversationId string) ([]string, error)
}

type ChatRepository struct {
	db func() *sql.DB
}

func NewChatRepository(db func() *sql.DB) ChatRepositoryContract {
	return ChatRepository{db: db}
}

func (u ChatRepository) CheckUserInConversation(userId string, conversationId string) (bool, error) {
	var userIdExists string

	query := fmt.Sprintf(`SELECT user_id FROM users_conversations WHERE conversation_id=%s AND user_id=%s`, conversationId, userId)

	rows, err := u.db().Query(query)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&userIdExists)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return true, nil
}

func (u ChatRepository) CreateNewConverasation(firstUserID string, secondUserID string) (string, error) {
	var err error
	var stmt *sql.Stmt
	query := fmt.Sprintf(`INSERT INTO dbo.conversation DEFAULT VALUES; SELECT SCOPE_IDENTITY() AS conversation_id;`)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return "", err
	}
	defer stmt.Close()

	var conversationId string

	if err := u.db().QueryRow(query).Scan(&conversationId); err != nil {
		fmt.Println(err)
		return "", err
	}

	query = fmt.Sprintf(`INSERT INTO users_conversations(conversation_id, user_id) VALUES(%s, %s),(%s, %s);`, conversationId, firstUserID, conversationId, secondUserID)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return "", err
	}

	if _, err := u.db().Exec(query); err != nil {
		fmt.Println(err)
		return "", err
	}

	return conversationId, nil
}

func (u ChatRepository) GetAllChatsByUserId(userId string) ([]datastruct.ChatsLastMsgs, error) {
	var msgs []datastruct.ChatsLastMsgs

	query := fmt.Sprintf(`SELECT * FROM (
		SELECT ROW_NUMBER() OVER (PARTITION BY message.conversation_id ORDER BY message.created_at DESC) AS LastMessage, users.username, users.profile_image, message.conversation_id, message.message_id, message.from_user, message.text, message.created_at
		FROM users_conversations
		JOIN message
		ON users_conversations.conversation_id = message.conversation_id
		JOIN users on message.from_user = users.id
		WHERE users_conversations.user_id=%s) AS a WHERE a.LastMessage = 1;`, userId)

	rows, err := u.db().Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var msg datastruct.ChatsLastMsgs
		err = rows.Scan(&msg.LastMessage, &msg.Username, &msg.ProfileImage, &msg.ConversationID, &msg.MessageID, &msg.FromUser, &msg.Text, &msg.CreatedAt)

		if err != nil {
			fmt.Println(err)
			continue
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func (u ChatRepository) CreateMessage(from_user string, conversation_id string, text string) (int64, error) {
	var err error
	var stmt *sql.Stmt
	var id int64

	query := fmt.Sprintf(`INSERT INTO dbo.message(from_user, conversation_id, text) VALUES('%s', %s, '%s'); 
	SELECT SCOPE_IDENTITY()`, from_user, conversation_id, text)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return 0, err
	}

	if err := u.db().QueryRow(query).Scan(&id); err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer stmt.Close()

	return id, err
}

func (u ChatRepository) GetConversationMessages(conversationId string) ([]datastruct.ChatMessage, error) {
	var msgs []datastruct.ChatMessage

	query := fmt.Sprintf(`
	SELECT users.username,message.conversation_id, message.message_id, message.from_user, message.text, message.created_at
	FROM users_conversations
	JOIN message
	ON users_conversations.conversation_id = message.conversation_id
	JOIN users on message.from_user = users.id
	WHERE users_conversations.user_id=message.from_user and message.conversation_id=%s ORDER BY message.created_at ASC;`, conversationId)

	rows, err := u.db().Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var msg datastruct.ChatMessage
		err = rows.Scan(&msg.Username, &msg.ConversationId, &msg.MessageId, &msg.FromUser, &msg.Text, &msg.CreatedAt)

		if err != nil {
			fmt.Println(err)
			continue
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func (u ChatRepository) GetConversationMemebersIDs(conversationId string) ([]string, error) {
	var membersIDs []string

	query := fmt.Sprintf(`SELECT user_id FROM users_conversations WHERE conversation_id='%s'`, conversationId)

	rows, err := u.db().Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var memberId string
		err = rows.Scan(&memberId)

		if err != nil {
			fmt.Println(err)
			continue
		}

		membersIDs = append(membersIDs, memberId)
	}

	return membersIDs, nil
}
