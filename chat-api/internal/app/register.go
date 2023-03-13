package app

import (
	"chat-project-go/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Services) Register(c *gin.Context) {
	user := new(dto.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	createdUserID, _ := s.authService.SignUp(user)

	c.JSON(http.StatusOK, createdUserID)
}
