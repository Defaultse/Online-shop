package main

import (
	"chat-project-go/internal/app"
	"chat-project-go/internal/drivers/mssql"
	"chat-project-go/internal/repository"
	"chat-project-go/internal/service"
	"chat-project-go/pkg/websocket"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

func main() {
	userRepository := repository.NewUserRepository(mssql.Connect)
	chatRepository := repository.NewChatRepository(mssql.Connect)
	feedRepository := repository.NewFeedsRepository(mssql.Connect)
	friendshipRepository := repository.NewFriendshipRepository(mssql.Connect)

	jwtTokenService := service.NewTokenManager(signingKey)
	authService := service.NewAuthService(jwtTokenService, userRepository)
	chatService := service.NewChatService(chatRepository)
	feedService := service.NewFeedService(feedRepository)
	friendship := service.NewFriendsService(friendshipRepository, userRepository)
	profileImageService := service.NewProfileImageService(userRepository)

	pool := websocket.NewPool()

	services := app.NewServices(authService, chatService, feedService, friendship, jwtTokenService, profileImageService, pool)

	go pool.Start()

	router := gin.Default()
	router.GET("/ping/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	})

	router.GET("/feeds/", services.GetFeeds)
	router.POST("/feeds/", services.PostFeed)
	router.POST("/feeds/comment", nil)
	router.POST("/feeds/like", nil)
	router.POST("/feeds/dislike", nil)

	router.GET("/friends/search/", services.SearchFriend)
	router.POST("/friends/invite/", services.SendFriendshipInvite)
	router.GET("/friends/pending/", services.FriendshipPending)
	router.POST("/friends/accept/", services.AcceptInvite)

	router.GET("/profile-images/", services.GetProfileImage)
	router.POST("/profile-images/", services.UploadImage)
	router.GET("/img/:asset", func(c *gin.Context) {
		dir := "storage"
		asset := c.Param("asset")
		if strings.TrimPrefix(asset, "/") == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		fullName := filepath.Join(dir, filepath.FromSlash(path.Clean("/"+asset)))
		c.File(fullName)
	})

	router.POST("/register/", services.Register)
	router.POST("/login/", services.Login)
	router.GET("/ws/chats/", func(ctx *gin.Context) {
		services.ServeWs(pool, ctx)
	})

	router.GET("ws/test/", func(ctx *gin.Context) {
		services.ServeTestWs(pool, ctx)
	})

	router.Run()
}
