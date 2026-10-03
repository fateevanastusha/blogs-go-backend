package model

type Blog struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	WebsiteURL  string `json:"websiteURL"`
}
