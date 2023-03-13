package dto

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"password_hash"`
	ProfileImage string `json:"profile_image"`
}

type UserType string

const (
	REGULAR UserType = "regular"
	VIP     UserType = "vip"
)
