package service

import (
	"chat-project-go/internal/dto"
	"chat-project-go/internal/repository"
	"crypto/md5"
	"fmt"
	"strconv"
)

type AuthService interface {
	SignUp(user *dto.User) (*int64, error)
	SignIn(email, password string) (*dto.User, *string, error)
	Logout(userID int64) error
}

type authService struct {
	userRepository repository.UserRepositoryContract
	tokenManager   TokenManager
}

func NewAuthService(tokenManager TokenManager, userRepository repository.UserRepositoryContract) AuthService {
	return &authService{
		userRepository: userRepository,
		tokenManager:   tokenManager,
	}
}

func (a *authService) SignUp(user *dto.User) (*int64, error) {
	user.PasswordHash = fmt.Sprintf("%x", md5.Sum([]byte(user.PasswordHash)))

	id, err := a.userRepository.CreateUser(*user)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (a *authService) SignIn(email, reqPassword string) (*dto.User, *string, error) {
	reqPasswordHash := fmt.Sprintf("%x", md5.Sum([]byte(reqPassword)))

	result, err := a.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, nil, err
	}

	user := dto.User{
		Id:           result.Id,
		Username:     result.Username,
		Name:         result.Name,
		Surname:      result.Surname,
		Email:        result.Email,
		Phone:        result.Phone,
		PasswordHash: result.PasswordHash,
		ProfileImage: result.ProfileImage,
	}

	if user.PasswordHash != reqPasswordHash {
		return nil, nil, fmt.Errorf("passwords don't match")
	} else {
		jwt, err := a.tokenManager.NewJWT(strconv.Itoa(int(user.Id)))
		if err != nil {
			return nil, nil, err
		}

		return &user, &jwt, nil
	}
}

func (a *authService) Logout(userID int64) error {
	_, err := a.tokenManager.NewJWT(strconv.Itoa(int(userID)))
	if err != nil {
		return err
	}
	return nil
}
