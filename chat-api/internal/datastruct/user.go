package datastruct

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"password_hash"`
	UserType     string `db:"role"`
	ProfileImage string `json:"profile_image"`
}
