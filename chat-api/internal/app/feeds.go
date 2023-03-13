package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Services) GetFeeds(c *gin.Context) {
	fmt.Println("Not implemented")
}

func (s *Services) PostFeed(c *gin.Context) {
	fmt.Println("Not implemented")
	// feed := new(dto.Feed)

	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// }

	// createdUserID, _ := s.authService.SignUp(user)

	// c.JSON(http.StatusOK, createdUserID)
}
