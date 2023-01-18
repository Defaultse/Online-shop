package app

import (
	"chat-project-go/internal/service"
	"chat-project-go/pkg/websocket"
)

type Services struct {
	authService     service.AuthService
	chatService     service.ChatService
	feedsService    service.FeedService
	friendService   service.FriendService
	profileService  service.ProfileImageService
	jwtTokenService service.TokenManager
	wsPool          *websocket.Pool
}

func NewServices(
	authService service.AuthService,
	chatService service.ChatService,
	feedsService service.FeedService,
	friendService service.FriendService,
	jwtTokenService service.TokenManager,
	profileImageService service.ProfileImageService,
	wsPool *websocket.Pool,
) *Services {
	return &Services{
		authService:     authService,
		chatService:     chatService,
		feedsService:    feedsService,
		friendService:   friendService,
		profileService:  profileImageService,
		jwtTokenService: jwtTokenService,
		wsPool:          wsPool,
	}
}
