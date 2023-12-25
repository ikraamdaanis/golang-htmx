package types

type Form struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Published   bool   `json:"published"`
	Content     string `json:"content"`
	Description string `json:"description"`
	UserId      string `json:"user_id"`
	Views       int    `json:"views"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
