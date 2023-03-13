package repository

import (
	"chat-project-go/internal/datastruct"
	"chat-project-go/internal/dto"
	"database/sql"
	"fmt"
)

type UserRepositoryContract interface {
	CreateUser(user dto.User) (int64, error)
	GetUserByEmail(email string) (*datastruct.User, error)
	GetUserPasswordByEmail(email string) (string, error)
	GetUserIdByEmail(email string) (int64, error)
	Search(str string) (*[]datastruct.PossibleFriend, error)
	UpdateImage(str string, userID int64) error
}

type UserRepository struct {
	db func() *sql.DB
}

func NewUserRepository(db func() *sql.DB) UserRepositoryContract {
	return UserRepository{db: db}
}

func (u UserRepository) UpdateImage(str string, userID int64) error {
	fmt.Println(str, userID)
	query := fmt.Sprintf(`UPDATE dbo.users SET profile_image='%s' WHERE id='%d'`, str, userID)

	if _, err := u.db().Exec(query); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (u UserRepository) CreateUser(user dto.User) (int64, error) {
	var id int64
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(`INSERT INTO dbo.users (username, name, surname, email, phone, password_hash, role) 
		VALUES ('%s', '%s', '%s', '%s', '%v', '%s', '%s'); 
		SELECT SCOPE_IDENTITY()`,
		user.Username, user.Name, user.Surname, user.Email, user.Phone, user.PasswordHash, dto.REGULAR)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return 0, err
	}

	if err := u.db().QueryRow(query).Scan(&id); err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer stmt.Close()

	return id, nil
}

func (u UserRepository) GetUserByEmail(email string) (*datastruct.User, error) {
	var user datastruct.User
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(`SELECT * FROM dbo.users WHERE email='%s'`, email)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer stmt.Close()

	if err := u.db().QueryRow(query).Scan(&user.Id, &user.Username, &user.Name, &user.Surname, &user.Email, &user.Phone, &user.PasswordHash, &user.UserType, &user.ProfileImage); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) GetUserPasswordByEmail(email string) (string, error) {
	var passwordHash string
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(`SELECT password_hash FROM dbo.users WHERE email='%s'`, email)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return "", err
	}

	defer stmt.Close()

	if err := u.db().QueryRow(query).Scan(&passwordHash); err != nil {
		fmt.Println(err)
		return "", err
	}

	return passwordHash, nil
}

func (u UserRepository) GetUserIdByEmail(email string) (int64, error) {
	var id int64
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(`SELECT id FROM dbo.users WHERE email='%s'`, email)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer stmt.Close()

	if err := u.db().QueryRow(query).Scan(&id); err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (u UserRepository) Search(searchStr string) (*[]datastruct.PossibleFriend, error) {
	var users []datastruct.PossibleFriend

	query := `SELECT id, username, name, surname FROM users WHERE (username + '|' + name + '|' + surname) LIKE '%` + searchStr + `%';`

	rows, err := u.db().Query(query)

	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user datastruct.PossibleFriend
		err = rows.Scan(&user.Id, &user.Username, &user.Name, &user.Surname)

		if err != nil {
			fmt.Println(err)
			continue
		}

		users = append(users, user)
	}

	return &users, nil
}
