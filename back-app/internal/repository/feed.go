package repository

import (
	"chat-project-go/internal/dto"
	"database/sql"
	"fmt"
	"time"
)

type FeedRepositoryContract interface {
	Create(post dto.Feed) error
}

type FeedRepository struct {
	db func() *sql.DB
}

func NewFeedsRepository(db func() *sql.DB) FeedRepositoryContract {
	return FeedRepository{db: db}
}

func (f FeedRepository) Create(post dto.Feed) error {
	var id int64
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(`INSERT INTO dbo.posts (user_id, description, media_url, created_at) 
		VALUES ('%s', '%s', '%s', '%v');`, post.UserId, post.Description, post.MediaUrl, time.Now())

	if stmt, err = f.db().Prepare(query); err != nil {
		fmt.Println(err)
		return err
	}

	if err := f.db().QueryRow(query).Scan(&id); err != nil {
		fmt.Println(err)
		return err
	}

	defer stmt.Close()

	return nil
}
