package service

import (
	"chat-project-go/internal/datastruct"
	"chat-project-go/internal/repository"
	"fmt"
)

type FriendService interface {
	Search(searchStr string) (*[]datastruct.PossibleFriend, error)
	SendInvite(userID1, userID2 string) error
	GetPending(userID1 string) ([]datastruct.PendingInvite, error)
	AcceptInvite(id string)
}

type friendService struct {
	friendsRepo repository.FriendshipRepositoryContract
	usersRepo   repository.UserRepositoryContract
}

func NewFriendsService(
	friendsRepo repository.FriendshipRepositoryContract,
	usersRepo repository.UserRepositoryContract,
) FriendService {
	return &friendService{
		friendsRepo: friendsRepo,
		usersRepo:   usersRepo,
	}
}

func (f *friendService) Search(searchStr string) (*[]datastruct.PossibleFriend, error) {
	result, err := f.usersRepo.Search(searchStr)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (f *friendService) SendInvite(userID1, userID2 string) error {
	err := f.friendsRepo.CreatePending(userID1, userID2)

	if err != nil {
		return err
	}

	return nil
}

func (f *friendService) GetPending(userID string) ([]datastruct.PendingInvite, error) {
	result, err := f.friendsRepo.GetPending(userID)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (f *friendService) AcceptInvite(id string) {
	err := f.friendsRepo.AcceptPending(id)

	if err != nil {
		fmt.Println(err)
	}
}
