package model

type Post struct {
	ID               int    `json:"ID"`
	BlogID           int    `json:"blogID"`
	Title            string `json:"title"`
	ShortDescription string `json:"shortDescription"`
	Content          string `json:"content"`
}
