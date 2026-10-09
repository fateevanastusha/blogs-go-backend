package e2e

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
	"github.com/stretchr/testify/require"
)

/*BLOGS*/

func createBlog(t *testing.T, h http.Handler) model.Blog {
	rec := DoRequest(t, h, http.MethodPost, "/blogs", map[string]string{"name": "test blog", "description": "test description", "websiteURL": "https://x.com"})
	require.Equal(t, http.StatusCreated, rec.Code)
	var blog model.Blog
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blog))
	return blog
}

func getBlog(t *testing.T, h http.Handler, id int) model.Blog {
	rec := DoRequest(t, h, http.MethodGet, "/blogs/"+strconv.Itoa(id), nil)
	var blog model.Blog
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blog))
	return blog
}

func deleteBlog(t *testing.T, h http.Handler, id int) {
	rec := DoRequest(t, h, http.MethodDelete, "/blogs/"+strconv.Itoa(id), nil)
	require.Equal(t, http.StatusNoContent, rec.Code)
}

func getAllEmptyBlogsCheck(t *testing.T, h http.Handler, _ *httptest.ResponseRecorder) {
	rec := DoRequest(t, h, http.MethodGet, "/blogs", nil)
	require.Equal(t, http.StatusOK, rec.Code)
	var blogs []model.Blog
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blogs))
	require.Equal(t, len(blogs), 0)
}

/*POSTS*/

func getPost(t *testing.T, h http.Handler, id int) model.Post {
	rec := DoRequest(t, h, http.MethodGet, "/posts/"+strconv.Itoa(id), nil)
	var post model.Post
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &post))
	return post
}

func deletePost(t *testing.T, h http.Handler, id int) {
	rec := DoRequest(t, h, http.MethodDelete, "/posts/"+strconv.Itoa(id), nil)
	require.Equal(t, http.StatusNoContent, rec.Code)
}

func createPost(t *testing.T, h http.Handler, blogID int) model.Post {
	rec := DoRequest(t, h, http.MethodPost, "/posts", map[string]any{"title": "test post", "shortDescription": "test description", "content": "test content", "blogID": blogID})
	require.Equal(t, http.StatusCreated, rec.Code)
	var post model.Post
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &post))
	return post
}

func getAllEmptyPostsCheck(t *testing.T, h http.Handler, _ *httptest.ResponseRecorder) {
	rec := DoRequest(t, h, http.MethodGet, "/posts", nil)
	require.Equal(t, http.StatusOK, rec.Code)
	var posts []model.Post
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &posts))
	require.Equal(t, 0, len(posts))
}
