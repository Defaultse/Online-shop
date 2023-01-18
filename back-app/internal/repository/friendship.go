package repository

import (
	"chat-project-go/internal/datastruct"
	"database/sql"
	"fmt"
)

type FriendshipRepositoryContract interface {
	CreatePending(userID1, userID2 string) error
	GetPending(userID string) ([]datastruct.PendingInvite, error)
	AcceptPending(pendingID string) error
}

type FriendshipRepository struct {
	db func() *sql.DB
}

func NewFriendshipRepository(db func() *sql.DB) FriendshipRepositoryContract {
	return FriendshipRepository{db: db}
}

func (f FriendshipRepository) CreatePending(userID1, userID2 string) error {
	query := fmt.Sprintf("INSERT INTO pending_friend(fromUser, toUser) VALUES(%s, %s)", userID1, userID2)

	if _, err := f.db().Exec(query); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (f FriendshipRepository) GetPending(userID string) ([]datastruct.PendingInvite, error) {
	var pendindInvites []datastruct.PendingInvite

	query := fmt.Sprintf(`SELECT * FROM pending_friend WHERE toUser='%s'`, userID)

	rows, err := f.db().Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var invite datastruct.PendingInvite
		err = rows.Scan(&invite.Id, &invite.FromUser, &invite.ToUser)

		if err != nil {
			fmt.Println(err)
			continue
		}

		pendindInvites = append(pendindInvites, invite)
	}

	return pendindInvites, nil
}

func (f FriendshipRepository) AcceptPending(pendingID string) error {
	panic("Not implemented")

	query := fmt.Sprintf(`BEGIN TRANSACTION; 
						DELETE FROM pending_friend WHERE pendingID='%s';
						INSERT INTO friends(userID1, userID2) VALUES();
						COMMIT;`, pendingID)

	if _, err := f.db().Exec(query); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
