package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Services) GetProfileImage(c *gin.Context) {
	location := c.Query("imagePath")

	_, _, _, _, y := s.profileService.GetFromPath(location)

	// _, _, _, loadedImage := s.profileService.GetFromPath(location)

	// w.Header().Set("Content-Type", "image/png")
	c.Data(http.StatusOK, "application/octet-stream", y)
	// c.JSON(http.StatusOK, y)
}

func (s *Services) UploadImage(c *gin.Context) {
	authToken := c.GetHeader("AuthToken")

	fmt.Println(authToken)

	userId, err := s.jwtTokenService.Parse(authToken)

	if err != nil {
		fmt.Println(err)
	}

	file, header, err := c.Request.FormFile("upload")
	filename := header.Filename
	fmt.Println(header.Filename)

	out, err := os.Create("./storage/" + filename + ".png")

	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
		log.Fatal(err)
	}

	err = s.profileService.UpdateProfileImage(filename+".png", *userId)

	if err != nil {
		log.Fatal(err)
	}
}
