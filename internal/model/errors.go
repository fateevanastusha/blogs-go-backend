package model

import "errors"

var ErrNotFound = errors.New("not found")
var ErrBlogNotFound = errors.New("blog not found")
var ErrPostsExists = errors.New("exists posts for this blog")

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Fields []FieldError
}

func (e *ValidationError) Error() string { return "validation failed" }
