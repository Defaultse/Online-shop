package service

import (
	"chat-project-go/internal/repository"
	"fmt"
	"image"
	"image/png"
	"os"

	draw "golang.org/x/image/draw"
)

type ProfileImageService interface {
	GetFromPath(location string) ([]byte, image.Image, string, image.Image, []byte)
	UpdateProfileImage(filename string, userId int64) error
}

type profileImageService struct {
	userRepository repository.UserRepositoryContract
	tokenManager   TokenManager
}

func NewProfileImageService(usersRepo repository.UserRepositoryContract) ProfileImageService {
	return &profileImageService{
		userRepository: usersRepo,
	}
}

func (s *profileImageService) UpdateProfileImage(filename string, userId int64) error {
	err := s.userRepository.UpdateImage(filename, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *profileImageService) GetFromPath(location string) ([]byte, image.Image, string, image.Image, []byte) {
	existingImageFile, err := os.Open("../storage/2b.png")
	if err != nil {
		fmt.Println(err)
	}

	defer existingImageFile.Close()

	b, err := os.ReadFile("../storage/2b.png") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	// str := string(b)

	imageData, imageType, err := image.Decode(existingImageFile)
	if err != nil {
		fmt.Println(err)
	}

	existingImageFile.Seek(0, 0)

	loadedImage, err := png.Decode(existingImageFile)

	if err != nil {
		fmt.Println(err)
	}

	dst := image.NewRGBA(image.Rect(0, 0, loadedImage.Bounds().Max.X/2, loadedImage.Bounds().Max.Y/2))
	draw.NearestNeighbor.Scale(dst, dst.Rect, loadedImage, loadedImage.Bounds(), draw.Over, nil)

	//save the imgByte to file
	out, err := os.Create("./test/test.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, loadedImage)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 1111111111
	out2, err := os.Create("./test/compressed.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out2, dst)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	asd, err := os.ReadFile("./test/compressed.png")
	if err != nil {
		fmt.Print(err)
	}

	return b, imageData, imageType, loadedImage, asd
}
