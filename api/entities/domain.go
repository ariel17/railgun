package entities

// Domain represents the web site domain to be tested.
type Domain struct {
	ID int64 `json:"id"`
	UserID string `json:"user_id"`
	URL string `json:"url"`
	Code string `json:"code"`
	Verified bool `json:"verified"`
}
