package entities

type Domain struct {
	UserID string `json:"user_id"`
	URL string `json:"url"`
	Code string `json:"code"`
	Verified bool `json:"verified"`
}
