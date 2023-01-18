package app

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Services) SearchFriend(c *gin.Context) {
	searchStr, exist := c.GetQuery("searchString")

	if exist == false {
		c.String(http.StatusBadRequest, "Field not provided")
	}

	userList, err := s.friendService.Search(searchStr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userList)
	c.JSON(http.StatusOK, userList)
}

func (s *Services) SendFriendshipInvite(c *gin.Context) {
	authToken := c.GetHeader("AuthToken")

	userId, err := s.jwtTokenService.Parse(authToken)

	if err != nil {
		fmt.Println(err)
	}

	userIdStr := strconv.FormatInt(int64(*userId), 10)

	var body map[string]interface{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	toUserID := body["toUserID"].(string)

	if userIdStr == toUserID {
		fmt.Println(errors.New("Pending friendship invite to itself"))
	}

	err = s.friendService.SendInvite(userIdStr, toUserID)

	if err != nil {
		fmt.Println(err)
	}
}

func (s *Services) FriendshipPending(c *gin.Context) {
	authToken := c.GetHeader("AuthToken")

	userId, err := s.jwtTokenService.Parse(authToken)

	if err != nil {
		fmt.Println(err)
	}

	userIdStr := strconv.FormatInt(int64(*userId), 10)

	result, err := s.friendService.GetPending(userIdStr)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, result)
}

func (s *Services) AcceptInvite(c *gin.Context) {
	var body map[string]interface{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	inviteID := body["pendingInviteID"].(string)

	s.friendService.AcceptInvite(inviteID)
}
