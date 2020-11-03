package entities

type Domain struct {
	ID int `json:"id"`
	UserID string `json:"user_id"`
	URL string `json:"url"`
	Code string `json:"code"`
	Verified bool `json:"verified"`
}
