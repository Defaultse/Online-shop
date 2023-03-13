package dto

type Feed struct {
	UserId      string `json:"user_id"`
	Description string `json:"description"`
	MediaUrl    []byte `json:"media_url"`
}
