package e2e

import (
	"encoding/json"
	"maps"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
	"github.com/fateevanastusha/blogs-go-backend/internal/server"
	"github.com/stretchr/testify/require"
)

// existingBlog — плейсхолдер для blogID, в цикле подменяется на id созданного в тесте блога
const existingBlog = "existing"

func withBlogID(body map[string]any, blogID int) map[string]any {
	res := maps.Clone(body)
	if res["blogID"] == existingBlog {
		res["blogID"] = blogID
	}
	return res
}

func TestCreatePosts(t *testing.T) {

	cases := []postCase{
		{
			body:         map[string]any{"title": "test post", "shortDescription": "test short description", "content": "test content", "blogID": existingBlog},
			description:  "success case",
			expectStatus: http.StatusCreated,
			after: func(t *testing.T, s http.Handler, rec *httptest.ResponseRecorder) {
				var post model.Post
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &post))
				require.Positive(t, post.ID)
				require.Positive(t, post.BlogID)
				require.Equal(t, "test post", post.Title)
				require.Equal(t, "test short description", post.ShortDescription)
				require.Equal(t, "test content", post.Content)

				rec = DoRequest(t, s, http.MethodGet, "/posts/"+strconv.Itoa(post.ID), nil)
				require.Equal(t, http.StatusOK, rec.Code)
			},
		},
		{
			body:         map[string]any{"shortDescription": "d", "content": "c", "blogID": existingBlog},
			description:  "no title",
			expectStatus: http.StatusBadRequest,
			after:        getAllEmptyPostsCheck,
		},
		{
			body:         map[string]any{"title": "lalalalalalalalalalalalalaalalalala", "shortDescription": "d", "content": "c", "blogID": existingBlog},
			description:  "too long title",
			expectStatus: http.StatusBadRequest,
			after:        getAllEmptyPostsCheck,
		},
		{
			body:         map[string]any{"title": "lala", "shortDescription": "d", "content": "c"},
			description:  "no blogID",
			expectStatus: http.StatusBadRequest,
			after:        getAllEmptyPostsCheck,
		},
		{
			body:         map[string]any{"title": "lala", "shortDescription": "d", "content": "c", "blogID": 434444},
			description:  "not existing blog",
			expectStatus: http.StatusNotFound,
			after:        getAllEmptyPostsCheck,
		},
		{
			body:         map[string]any{},
			description:  "empty body",
			expectStatus: http.StatusBadRequest,
			after:        getAllEmptyPostsCheck,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			s, err := server.New()
			require.NoError(t, err)

			blog := createBlog(t, s)

			rec := DoRequest(t, s, http.MethodPost, "/posts", withBlogID(c.body, blog.ID))
			require.Equal(t, c.expectStatus, rec.Code)
			if c.after != nil {
				c.after(t, s, rec)
			}

		})
	}
}

func TestPutPosts(t *testing.T) {
	cases := []postCase{
		{
			body:         map[string]any{"title": "edited", "shortDescription": "edited", "content": "edited"},
			description:  "success case",
			expectStatus: http.StatusCreated,
		},
		{
			body:         map[string]any{"shortDescription": "d", "content": "c"},
			description:  "no title",
			expectStatus: http.StatusBadRequest,
		},
		{
			body:         map[string]any{"title": "lalalalalalalalalalalalalaalalalala", "shortDescription": "d", "content": "c"},
			description:  "too long title",
			expectStatus: http.StatusBadRequest,
		},
		{
			body:         map[string]any{},
			description:  "empty body",
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {

			s, err := server.New()
			require.NoError(t, err)

			blog := createBlog(t, s)
			newPost := createPost(t, s, blog.ID)

			rec := DoRequest(t, s, http.MethodPut, "/posts/"+strconv.Itoa(newPost.ID), c.body)
			require.Equal(t, c.expectStatus, rec.Code)
			if c.expectStatus == http.StatusBadRequest {
				notEditedPost := getPost(t, s, newPost.ID)
				require.Equal(t, newPost.Title, notEditedPost.Title)
				require.Equal(t, newPost.ShortDescription, notEditedPost.ShortDescription)
				require.Equal(t, newPost.Content, notEditedPost.Content)
			}
			if c.expectStatus == http.StatusCreated {
				editedPost := getPost(t, s, newPost.ID)
				require.Equal(t, c.body["title"], editedPost.Title)
				require.Equal(t, c.body["shortDescription"], editedPost.ShortDescription)
				require.Equal(t, c.body["content"], editedPost.Content)
			}

			if c.after != nil {
				c.after(t, s, rec)
			}

		})
	}

}

func TestDeletePosts(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		blog := createBlog(t, s)
		newPost := createPost(t, s, blog.ID)

		rec := DoRequest(t, s, http.MethodDelete, "/posts/"+strconv.Itoa(newPost.ID), nil)
		require.Equal(t, http.StatusNoContent, rec.Code)
	})
	t.Run("delete already deleted", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		blog := createBlog(t, s)
		newPost := createPost(t, s, blog.ID)

		rec := DoRequest(t, s, http.MethodDelete, "/posts/"+strconv.Itoa(newPost.ID), nil)
		require.Equal(t, http.StatusNoContent, rec.Code)

		rec = DoRequest(t, s, http.MethodDelete, "/posts/"+strconv.Itoa(newPost.ID), nil)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("delete not existing", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		rec := DoRequest(t, s, http.MethodDelete, "/posts/434444", nil)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestGetByIdPosts(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		blog := createBlog(t, s)
		newPost := createPost(t, s, blog.ID)

		rec := DoRequest(t, s, http.MethodGet, "/posts/"+strconv.Itoa(newPost.ID), nil)
		require.Equal(t, http.StatusOK, rec.Code)
		var post model.Post
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &post))
		require.Equal(t, newPost.Title, post.Title)
		require.Equal(t, newPost.ShortDescription, post.ShortDescription)
		require.Equal(t, newPost.Content, post.Content)
		require.Equal(t, newPost.BlogID, post.BlogID)
	})

	t.Run("get not existing", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		rec := DoRequest(t, s, http.MethodGet, "/posts/434444", nil)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("get deleted", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)
		blog := createBlog(t, s)
		newPost := createPost(t, s, blog.ID)
		deletePost(t, s, newPost.ID)

		rec := DoRequest(t, s, http.MethodGet, "/posts/"+strconv.Itoa(newPost.ID), nil)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

}

func TestGetPosts(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		var posts []model.Post

		rec := DoRequest(t, s, http.MethodGet, "/posts", nil)
		require.Equal(t, http.StatusOK, rec.Code)
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &posts))
		require.Equal(t, 0, len(posts))

		blog := createBlog(t, s)
		newPosts := []model.Post{createPost(t, s, blog.ID), createPost(t, s, blog.ID), createPost(t, s, blog.ID)}

		rec = DoRequest(t, s, http.MethodGet, "/posts", nil)
		require.Equal(t, http.StatusOK, rec.Code)
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &posts))
		require.Equal(t, len(newPosts), len(posts))

		for i := 0; i < len(newPosts); i++ {
			p1, p2 := newPosts[i], posts[i]
			require.Equal(t, p1.Title, p2.Title)
			require.Equal(t, p1.ShortDescription, p2.ShortDescription)
			require.Equal(t, p1.Content, p2.Content)
			require.Equal(t, p1.BlogID, p2.BlogID)
		}

		for i := 0; i < len(newPosts); i++ {
			deletePost(t, s, newPosts[i].ID)
		}

		rec = DoRequest(t, s, http.MethodGet, "/posts", nil)
		require.Equal(t, http.StatusOK, rec.Code)
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &posts))
		require.Equal(t, 0, len(posts))
	})

}
