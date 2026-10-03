package model

type Post struct {
	ID               int    `json:"id"`
	BlogID           int    `json:"blogId"`
	Title            string `json:"title"`
	ShortDescription string `json:"shortDescription"`
	Content          string `json:"content"`
}
