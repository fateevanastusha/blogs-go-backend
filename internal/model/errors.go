package model

import "errors"

var ErrNotFound = errors.New("not found")
var ErrBlogNotFound = errors.New("blog not found")
var ErrPostsExists = errors.New("exists posts for this blog")
