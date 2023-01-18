package service

import "chat-project-go/internal/repository"

type FeedService interface {
}

type feedService struct {
	feedRepo *repository.FeedRepositoryContract
}

func NewFeedService(feedRepo repository.FeedRepositoryContract) FeedService {
	return &feedService{
		feedRepo: &feedRepo,
	}
}

func (f *feedService) PostFeed() {

}
